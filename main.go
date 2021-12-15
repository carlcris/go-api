package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Go Programming")
	})
	router.GET("/say-hello", SayHello)

	router.Run(":" + port)
}
