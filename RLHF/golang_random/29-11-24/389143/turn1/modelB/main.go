package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define our routes here
	v1 := router.Group("/v1")
	{
		v1.GET("/hello", helloHandler)
	}

	// Start the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
	}
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
