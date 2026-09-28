// server/main.go
package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// 中间件：日志
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${method} ${uri} ${status} ${latency_human}` + "\n",
	}))
	e.Use(middleware.Recover())

	// 捕获所有路径（支持 /任意uuid）
	e.Any("/*", func(c echo.Context) error {
		url := c.Request().URL.String()
		log.Printf("📥 收到请求: %s | RemoteAddr: %s", url, c.RealIP())

		return c.String(http.StatusOK, "OK")
	})

	port := getServerPortFromEnv()
	log.Printf("🚀 Echo 服务端已启动: http://localhost:%d", port)
	if err := e.Start(":" + strconv.Itoa(port)); err != nil {
		log.Fatal(err)
	}
}

func getServerPortFromEnv() int {
	const defaultPort = 10086
	portStr := os.Getenv("SERVER_PORT")
	if portStr == "" {
		log.Printf("⚠️  未设置 SERVER_PORT 环境变量，使用默认端口: %d", defaultPort)
		return defaultPort
	}

	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 || port > 65535 {
		log.Printf("⚠️  SERVER_PORT 格式错误或端口范围不正确，使用默认端口: %d", defaultPort)
		return defaultPort
	}

	return port
}

