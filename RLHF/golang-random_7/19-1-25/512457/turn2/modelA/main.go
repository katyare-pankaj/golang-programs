package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User represents the structure of our resource
type User struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required,min=3,max=100"`
}

// ApiResponse represents the consistent structure of all API responses
type ApiResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

// A slice to store users as a simulated database
var users = []User{}

// createUser handles the creation of a user
func createUser(c *gin.Context) {
	var newUser User

	// Bind JSON and validate input
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Status: "failure",
			Error:  err.Error(),
		})
		return
	}

	// Simulate setting a new user ID
	newUser.ID = len(users) + 1

	// Store the new user in the list
	users = append(users, newUser)

	// Respond with success
	c.JSON(http.StatusCreated, ApiResponse{
		Status: "success",
		Data:   newUser,
	})
}

func main() {
	r := gin.Default()

	// Register routes
	r.POST("/users", createUser)

	// Start the server
	r.Run(":8080")
}
