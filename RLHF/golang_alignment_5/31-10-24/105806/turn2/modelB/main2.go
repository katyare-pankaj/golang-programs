package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type SensorData struct {
	DeviceID    string  `json:"device_id"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Timestamp   string  `json:"timestamp"`
}

type AggregatedData struct {
	DeviceID string  `json:"device_id"`
	Count    int     `json:"count"`
	Sum      float32 `json:"sum"`
	Average  float32 `json:"average"`
}

var (
	aggregatedData map[string]*AggregatedData
	mux            sync.Mutex
)

func main() {
	aggregatedData = make(map[string]*AggregatedData)
	// Create a new Gin router
	router := gin.Default()

	// Endpoint to receive data from the device and aggregate it
	router.POST("/aggregate", func(c *gin.Context) {
		// Read the request body to get the data from the device
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading request body: %v", err))
			return
		}

		var data SensorData
		if err := json.Unmarshal(body, &data); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Invalid data format: %v", err))
			return
		}
		// Aggregate the data
		aggregateData(data)
		c.String(http.StatusOK, "Data aggregated successfully")
	})

	// Run the server on port 8081
	router.Run(":8081")
}

func aggregateData(data SensorData) {
	mux.Lock()
	defer mux.Unlock()

	deviceID := data.DeviceID
	temp := data.Temperature

	if _, ok := aggregatedData[deviceID]; !ok {
		aggregatedData[deviceID] = &AggregatedData{
			DeviceID: deviceID,
			Count:    1,
			Sum:      temp,
			Average:  temp,
		}
	} else {
		agg := aggregatedData[deviceID]
		agg.Count++
		agg.Sum += temp
		agg.Average = agg.Sum / float32(agg.Count)
	}
}
