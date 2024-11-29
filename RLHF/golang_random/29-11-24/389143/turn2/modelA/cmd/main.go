package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/29-11-24/389143/turn2/modelA/handlers"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

func main() {
	router := gin.Default()

	// Create an instance of the HelloHandler
	h := handlers.NewHelloHandler()

	// Define the routes
	v1 := router.Group("/v1")
	{
		v1.GET("/hello", h.HandleHello)
	}

	// Start the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
	}
}
