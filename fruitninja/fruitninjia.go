package fruitninja

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func FruitNinjaRouterSetup() *gin.Engine {
	r := gin.Default()

	r.GET("/:fruit/:count", func(ctx *gin.Context) {
		var msg string

		fruit := ctx.Param("fruit")
		cnt, err := strconv.Atoi(ctx.Param("count"))
		if err != nil {
			ctx.String(500, fmt.Sprintf("%s: %s", "🐞", err.Error()))
		}

		switch fruit {
		case "apple":
			msg = strings.Repeat("🍎", cnt)
		case "banana":
			msg = strings.Repeat("🍌", cnt)
		case "orange":
			msg = strings.Repeat("🍊", cnt)
		case "watermelon":
			msg = strings.Repeat("🍉", cnt)
		case "pear":
			msg = strings.Repeat("🍐", cnt)
		case "cherry":
			msg = strings.Repeat("🍒", cnt)
		case "strawberry":
			msg = strings.Repeat("🍓", cnt)
		case "kiwi":
			msg = strings.Repeat("🥝", cnt)
		default:
			msg = "🐞"
		}

		ctx.String(200, msg)
	})

	return r
}
