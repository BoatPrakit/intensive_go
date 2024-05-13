package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.Run(":" + os.Getenv("APP_PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
