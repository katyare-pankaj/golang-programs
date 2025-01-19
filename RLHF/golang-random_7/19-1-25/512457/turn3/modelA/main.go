package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User represents a user in the system
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required,min=3,max=100"`
}

// ApiResponse represents the structure of the API response
type ApiResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

// In-memory user store
var users = []User{}

// getNextUserID generates the next user ID
func getNextUserID() int {
	return len(users) + 1
}

// getUsers retrieves the list of users
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, ApiResponse{
		Status: "success",
		Data:   users,
	})
}

// createUser creates a new user
func createUser(c *gin.Context) {
	var newUser User

	// Bind JSON with validation
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Status: "failure",
			Error:  err.Error(),
		})
		return
	}

	// Assign a new ID to the user
	newUser.ID = getNextUserID()

	// Add the user to the store
	users = append(users, newUser)

	// Return success response
	c.JSON(http.StatusCreated, ApiResponse{
		Status: "success",
		Data:   newUser,
	})
}

func main() {
	r := gin.Default()

	// Register routes
	r.GET("/users", getUsers)
	r.POST("/users", createUser)

	// Start the server
	r.Run(":8080")
}
