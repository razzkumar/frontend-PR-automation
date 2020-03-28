package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/razzkumar/sfebuild-tool/webhook"
)

//const secrect = "topsecrectpassword"
func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "4001"
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Name":    "ghWebHook",
			"version": "0.0.4",
		})
	})
	router.POST("/webhook", webhook.Handler())
	router.Run(":" + port)
}
