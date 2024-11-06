package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SensorData struct {
	DeviceID    string  `json:"device_id"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Timestamp   string  `json:"timestamp"`
}

func main() {
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a new Gin router
	router := gin.Default()

	// Endpoint to simulate device data generation
	router.GET("/data", func(c *gin.Context) {
		// Generate random data and send it back as a response
		data := SensorData{
			DeviceID:    "device123",
			Temperature: rand.Float32() * 20,
			Humidity:    rand.Float32() * 100,
			Timestamp:   time.Now().Format(time.RFC3339),
		}
		c.JSON(http.StatusOK, data)
	})

	// Run the server on port 8080
	router.Run(":8080")
}
