// client/main.go
package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func main() {
	serverURL := getServerURLFromEnv()
	interval := getIntervalFromEnv()
	warnTimeout := getWarnTimeoutFromEnv()

	log.Printf("🚀 HTTP 客户端启动")
	log.Printf("   服务端地址: %s", serverURL)
	log.Printf("   发送间隔: %d 秒", int(interval.Seconds()))
	log.Printf("   超时警告阈值: %d 秒", int(warnTimeout.Seconds()))

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		sendRequest(serverURL, warnTimeout)
	}
}

// 获取服务端地址（支持环境变量）
func getServerURLFromEnv() string {
	url := os.Getenv("SERVER_URL")
	if url == "" {
		url = "http://localhost:8080"
		log.Printf("⚠️  未设置 SERVER_URL 环境变量，使用默认值: %s", url)
	}

	// 移除末尾的斜杠，避免出现 //
	url = strings.TrimRight(url, "/")
	return url
}

// 获取发送间隔（单位：秒）
func getIntervalFromEnv() time.Duration {
	const defaultInterval = 5

	env := os.Getenv("INTERVAL")
	if env == "" {
		log.Printf("⚠️  未设置 INTERVAL 环境变量，使用默认值: %d 秒", defaultInterval)
		return defaultInterval * time.Second
	}

	seconds, err := strconv.Atoi(env)
	if err != nil || seconds <= 0 {
		log.Printf("⚠️  INTERVAL 格式错误，使用默认值: %d 秒", defaultInterval)
		return defaultInterval * time.Second
	}

	return time.Duration(seconds) * time.Second
}

// 获取超时警告阈值（单位：秒）
func getWarnTimeoutFromEnv() time.Duration {
	const defaultWarnTimeout = 5

	env := os.Getenv("WARN_TIMEOUT")
	if env == "" {
		log.Printf("⚠️  未设置 WARN_TIMEOUT 环境变量，使用默认值: %d 秒", defaultWarnTimeout)
		return defaultWarnTimeout * time.Second
	}

	seconds, err := strconv.Atoi(env)
	if err != nil || seconds <= 0 {
		log.Printf("⚠️  WARN_TIMEOUT 格式错误，使用默认值: %d 秒", defaultWarnTimeout)
		return defaultWarnTimeout * time.Second
	}

	return time.Duration(seconds) * time.Second
}

func sendRequest(serverURL string, warnTimeout time.Duration) {
	u := uuid.New()
	path := "/" + u.String()
	fullURL := serverURL + path

	startedAt := time.Now()
	log.Printf("📤 发送请求 [%s] UUID: %s | URL: %s", startedAt.Format("2006-01-02 15:04:05.000"), u.String(), fullURL)

	warnTimer := time.AfterFunc(warnTimeout, func() {
		log.Printf("⚠️ 请求已超过 %d 秒，仍在等待服务器响应: UUID=%s URL=%s", int(warnTimeout.Seconds()), u.String(), fullURL)
	})

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(fullURL)
	warnTimer.Stop()

	if err != nil {
		log.Printf("❌ 请求失败: %v", err)
		return
	}
	defer resp.Body.Close()

	elapsed := time.Since(startedAt)
	log.Printf("✅ 响应状态码: %d | 耗时: %s", resp.StatusCode, elapsed)
}

