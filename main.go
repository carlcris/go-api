package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"message": "Hello " + name})
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
	router.GET("/say-hello/:name", SayHello)

	router.Run(":" + port)
}
