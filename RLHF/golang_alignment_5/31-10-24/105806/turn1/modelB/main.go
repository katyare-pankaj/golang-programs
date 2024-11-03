package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a new Gin router
	router := gin.Default()

	// Endpoint to simulate device data generation
	router.GET("/data", func(c *gin.Context) {
		// Generate random data and send it back as a response
		data := rand.Intn(100)
		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// Run the server on port 8080
	router.Run(":8080")
}
