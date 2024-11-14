package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define a simple endpoint for retrieving optimization results
	router.GET("/optimization_results", func(c *gin.Context) {
		results := []map[string]interface{}{
			{
				"node":    "Node1",
				"latency": 12.3,
			},
			{
				"node":    "Node2",
				"latency": 8.9,
			},
		}
		c.JSON(200, results)
	})

	// Start the server
	router.Run(":8080")
}
