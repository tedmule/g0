package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/sijms/go-ora/v2"
)

// Message struct for DingTalk message
type Message struct {
	MsgType  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type At struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

func sendDingTalkMessage(title, content string, atMobiles []string, atAll bool) bool {
	// Get environment variables with defaults
	accessToken := getEnvOrDefault("DINGDING_ACCESS_TOKEN", "476fbffae8baa6fabdac40da712ed8825f86e787b3d2445305687ada5e151d60")
	secret := getEnvOrDefault("DINGDING_ACCESS_SECRET", "SECa3a4af3a2557a79c582c0e34785d673bf513cdaae4175a64f1f4d344b4f3d8d7")

	// Generate timestamp and signature
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, secret)

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	sign := url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))

	// Construct webhook URL
	webhookURL := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%s&sign=%s",
		accessToken, timestamp, sign)

	// Create message struct
	msg := Message{
		MsgType: "markdown",
		Markdown: Markdown{
			Title: title,
			Text:  content,
		},
		At: At{
			IsAtAll: atAll,
		},
	}

	if atMobiles != nil {
		msg.At.AtMobiles = atMobiles
	}

	// Convert to JSON
	jsonData, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("🐛 JSON marshal error: %v\n", err)
		return false
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("🐛 Request creation error: %v\n", err)
		return false
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("🐛 Request error: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("🐛 Request failed with status code: %d\n", resp.StatusCode)
		return false
	}

	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("🐛 Response read error: %v\n", err)
		return false
	}

	var result struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("🐛 Response parse error: %v\n", err)
		return false
	}

	if result.ErrCode == 0 {
		fmt.Println("🐛 消息发送成功")
		return true
	}

	fmt.Printf("🐛 消息发送失败: %s\n", result.ErrMsg)
	return false
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func dieOnError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s %v\n", msg, err)
		os.Exit(1)
	}
}

func main() {
	percentStr := getEnvOrDefault("ORACLE_TS_PERCENT", "80")
	percent, err := strconv.ParseFloat(percentStr, 64) // 64 表示转为 float64
	username := getEnvOrDefault("ORACLE_TS_USERNAME", "username")
	password := getEnvOrDefault("ORACLE_TS_USERNAME", "password")
	connStr := getEnvOrDefault("ORACLE_TS_CONN", "localhost:1521/name")
	if err != nil {
		fmt.Println("转换失败:", err)
		percent = 80
	}

	tsSQL := `
	SELECT
		t.tablespace_name,
		t.total_space,
		NVL(f.free_space, 0)
	FROM
		(SELECT tablespace_name, SUM(bytes) total_space
		 FROM dba_data_files
		 GROUP BY tablespace_name) t
	LEFT JOIN
		(SELECT tablespace_name, SUM(bytes) free_space
		 FROM dba_free_space
		 GROUP BY tablespace_name) f
	ON t.tablespace_name = f.tablespace_name
	`

	// Simplified DSN for go-ora/v2
	dsn := fmt.Sprintf("oracle://%s:%s@%s", username, password, connStr)

	conn, err := sql.Open("oracle", dsn)
	dieOnError("error while connecting to "+connStr+":", err)
	defer conn.Close()

	err = conn.Ping()
	dieOnError("Ping failed:", err)

	stmt, err := conn.Prepare(tsSQL)
	dieOnError("Can't prepare query:", err)
	defer stmt.Close()

	rows, err := stmt.Query()
	dieOnError("Can't create query:", err)
	defer rows.Close()

	var messages []string
	for rows.Next() {
		var tablespaceName string
		var totalSpace, freeSpace int64

		err := rows.Scan(&tablespaceName, &totalSpace, &freeSpace)
		if err != nil {
			fmt.Println("Scan error:", err)
			break
		}

		// Calculate usage percentage
		usedSpace := totalSpace - freeSpace
		usagePercent := float64(usedSpace) / float64(totalSpace) * 100

		fmt.Printf("Tablespace: %s, Total Space: %d, Free Space: %d, Usage: %.2f%%\n",
			tablespaceName, totalSpace, freeSpace, usagePercent)

		// Check if usage exceeds 80%
		if usagePercent > percent {
			// messages = append(messages, fmt.Sprintf("%s | %d/%d | **%.2f%%**", tablespaceName, freeSpace, totalSpace, usagePercent))
			messages = append(messages, fmt.Sprintf("%s: **%.2f%%**", tablespaceName, usagePercent))
		}
	}
	dingdingMessage := fmt.Sprintf("## 🔎 Oracle表空间检查\n\n %s", strings.Join(messages, "\n\n"))
	sendDingTalkMessage(
		"Oracle表空间检查",
		dingdingMessage,
		nil,   // atMobiles (can be customized)
		false, // atAll
	)
}
