package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nlopes/slack"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/say/:text/channel/:channel", func(c *gin.Context) {
		text := c.Params.ByName("text")
		channel := c.Params.ByName("channel")

		token := getEnv("SLACK_TOKEN", "WRONG_SLACK_TOKEN")
		api := slack.New(token)
		respChannel, respTimestamp, err := api.PostMessage(channel, text, slack.PostMessageParameters{})

		if err != nil {
			c.JSON(500, gin.H{
				"result": "FAILED",
				"error":  err,
			})
		} else {
			c.JSON(200, gin.H{
				"result":  "OK",
				"channel": respChannel,
				"time":    respTimestamp,
			})
		}
	})

	return r
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	r := setupRouter()
	r.Run(":" + getEnv("PORT", "3000"))
}
