// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SensorData represents the data structure for sensor data
type SensorData struct {
	DeviceID    string  `json:"device_id"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Timestamp   int64   `json:"timestamp"`
}

var (
	aggregatedData []SensorData
)

func main() {
	// Initialize router
	router := gin.Default()

	// Route for device to send data
	router.POST("/collect", collectSensorData)

	// Route to start data aggregation
	router.GET("/aggregate", aggregateData)

	// Run the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}

// collectSensorData handles POST requests from devices to collect sensor data
func collectSensorData(c *gin.Context) {
	var data SensorData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.Timestamp = time.Now().UnixNano()
	aggregatedData = append(aggregatedData, data)
	c.JSON(http.StatusOK, gin.H{"message": "Data collected successfully"})
}

// aggregateData handles GET requests to aggregate and report the sensor data
func aggregateData(c *gin.Context) {
	if len(aggregatedData) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No data available for aggregation"})
		return
	}

	// Perform aggregation (in this case, we'll just calculate the average)
	var totalTemp, totalHumidity float32
	for _, entry := range aggregatedData {
		totalTemp += entry.Temperature
		totalHumidity += entry.Humidity
	}

	averageTemp := totalTemp / float32(len(aggregatedData))
	averageHumidity := totalHumidity / float32(len(aggregatedData))

	aggregationReport := gin.H{
		"average_temperature": averageTemp,
		"average_humidity":    averageHumidity,
		"data_count":          len(aggregatedData),
	}

	// Reset aggregated data for next aggregation period
	aggregatedData = aggregatedData[:0]

	c.JSON(http.StatusOK, aggregationReport)
}
