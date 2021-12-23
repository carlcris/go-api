package main

import (
	"log"
	"net/http"
	"os"

	"go-api/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SayHello(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"message": "Hello " + name})
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Go Programming")
	})
	router.GET("/say-hello/:name", SayHello)
	router.GET("/patients", model.GetPatientList)
	router.GET("/patients/patient/:id", model.GetUserByID)

	router.Run(":" + port)
}
