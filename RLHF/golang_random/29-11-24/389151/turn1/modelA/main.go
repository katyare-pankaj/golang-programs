package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	// Simulate some checks, e.g., database connection, cache status
	// In a real application, replace these with actual checks
	c.JSON(http.StatusOK, gin.H{
		"message": "Application is healthy",
	})
}

func main() {
	r := gin.Default()

	// Register the health check endpoint
	r.GET("/health", healthCheck)

	// Start the server
	log.Fatal(r.Run(":8080"))
}
