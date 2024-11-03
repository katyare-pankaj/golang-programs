package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Endpoint to receive data from the device
	router.POST("/process", func(c *gin.Context) {
		// Read the request body to get the data from the device
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading request body: %v", err))
			return
		}

		// Process the received data
		fmt.Printf("Received data: %s\n", body)

		c.String(http.StatusOK, "Data processed successfully")
	})

	// Run the server on port 8081
	router.Run(":8081")
}
