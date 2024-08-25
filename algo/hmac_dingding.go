package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

func main() {
	// 获取当前时间戳（毫秒）
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)

	// 秘密密钥
	secret := "SECef5320c24ee3182284b33fb30cc8af1a437a1fbc95d61827e9cded61b790e7f8"
	secretEnc := []byte(secret)

	// 构建待签名的字符串
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	stringToSignEnc := []byte(stringToSign)

	// 生成 HMAC-SHA256 签名
	h := hmac.New(sha256.New, secretEnc)
	h.Write(stringToSignEnc)
	hmacCode := h.Sum(nil)

	// Base64 编码签名
	sign := base64.StdEncoding.EncodeToString(hmacCode)

	// URL 编码签名
	sign = url.QueryEscape(sign)

	// 打印时间戳和签名
	fmt.Println(timestamp)
	fmt.Println(sign)
}
