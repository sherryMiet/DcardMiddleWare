package main

import (
	"errors"
	"time"

	limiter "Ratelimit/limiter"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	lm := limiter.NewRateLimiter(time.Hour, 1000, func(ctx *gin.Context) (string, error) {
		key := ctx.Request.Header.Get("X-API-KEY")
		if key != "" {
			return key, nil
		}
		return "", errors.New("API key is missing")
	})

	r.GET("/DrawCard", lm.Middleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "抽卡",
		})
	})

	r.Run(":8080")
}
