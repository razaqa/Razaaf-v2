package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"

	"github.com/razaqa/razaaf-v2/config"
)

func main() {
	
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := config.GetAppConfig().Port
	router.Run(":" + port)
}