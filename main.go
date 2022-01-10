package main

import (
	"log"
	"net/http"
	"os"

	"go-api/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SayHello(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"message": "Hello " + name})
}
func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Patient Service Running")
}
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", welcome)
	router.GET("/say-hello/:name", SayHello)
	router.GET("/patients", model.GetPatientList)
	router.GET("/patients/:id", model.GetPatientByID)
	router.GET("/patients-add", model.GetPatientAddress)
	router.GET("/patients/address/:id", model.GetPatientAddressById)
	router.Run(":" + port)
}
