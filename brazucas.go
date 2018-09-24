package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	app := gin.New()

	app.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})

	app.Run(":" + os.Getenv("PORT"))
}
