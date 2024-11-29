package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

type ServiceStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// Example health check function for a database service
func checkDatabaseStatus() string {
	// Perform database connection check or any other health check logic here
	// For demonstration purposes, let's randomly return "up" or "down"
	if rand.Intn(2) == 0 {
		return "up"
	}
	return "down"
}

// Example health check function for a cache service
func checkCacheStatus() string {
	// Perform cache connection check here
	// Similar to the database check, let's return random status
	if rand.Intn(2) == 0 {
		return "up"
	}
	return "down"
}

func healthCheck(c *gin.Context) {
	services := []ServiceStatus{
		{
			Name:   "Database",
			Status: checkDatabaseStatus(),
		},
		{
			Name:   "Cache",
			Status: checkCacheStatus(),
		},
	}

	var allServicesUp bool = true
	for _, service := range services {
		if service.Status == "down" {
			allServicesUp = false
			break
		}
	}

	overallStatus := "up"
	if !allServicesUp {
		overallStatus = "down"
	}

	c.JSON(http.StatusOK, gin.H{
		"message":           fmt.Sprintf("Application health: %s", overallStatus),
		"dependentServices": services,
	})
}

func main() {
	r := gin.Default()
	r.GET("/health", healthCheck)
	log.Fatal(r.Run(":8080"))
}
