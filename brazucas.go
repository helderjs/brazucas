package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Handshake struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}

func main() {
	app := gin.New()

	app.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})

	app.POST("/events", func(context *gin.Context) {
		var payload Handshake

		context.ShouldBindJSON(&payload)

		context.String(http.StatusOK, payload.Challenge)
	})

	app.Run(":" + os.Getenv("PORT"))
}
