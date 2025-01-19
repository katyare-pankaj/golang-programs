package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define our ApiResponse struct
type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Function to create a new ApiResponse with appropriate status and message
func NewApiResponse(status string, message string, data interface{}, error string) ApiResponse {
	return ApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   error,
	}
}

// User represents the structure of our user resource
type User struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required,min=3,max=100"`
}

func main() {
	r := gin.Default()

	// Middleware to handle errors and apply consistent error handling
	r.Use(func(c *gin.Context) {
		c.Next()

		// Check if the value of 'ApiResponse' is set in the Gin context
		if res, ok := c.Get("ApiResponse"); ok {
			//Convert to ApiResponse type
			apiResponse := res.(ApiResponse)

			// If error is present in the response, send it back in the JSON format
			if apiResponse.Error != "" {
				c.JSON(http.StatusInternalServerError, apiResponse)
			} else {
				// Otherwise, send the response
				c.JSON(http.StatusOK, apiResponse)
			}
		}
	})

	// Set routing
	r.POST("/users", createUser)
	r.GET("/users/:id", getUser)

	// Start the server on port 8080
	r.Run(":8080")
}

func createUser(c *gin.Context) {
	var newUser User

	// Bind JSON and validate input
	if err := c.ShouldBindJSON(&newUser); err != nil {
		// Validatopn error, create an ApiResponse with an error and abort the request
		c.Set("ApiResponse", NewApiResponse("failure", "Validation error", nil, err.Error()))
		c.Abort()
		return
	}

	// Simulate database operations
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	// Set success ApiResponse and no need to call `c.Abort()`
	c.Set("ApiResponse", NewApiResponse("success", "User created successfully", newUser, ""))
}

func getUser(c *gin.Context) {
	userID, err := getUserIDFromParam(c.Param("id"))
	if err != nil {
		c.Set("ApiResponse", NewApiResponse("failure", "Invalid user ID", nil, err.Error()))
		c.Abort()
		return
	}

	user, found := findUserByID(userID)
	if !found {
		c.Set("ApiResponse", NewApiResponse("failure", "User not found", nil, "User not found"))
		c.Abort()
		return
	}

	c.Set("ApiResponse", NewApiResponse("success", "User found", user, ""))
}

// Helper function to get user id from the parameter
func getUserIDFromParam(idParam string) (int, error) {
	userID := 0
	err := fmt.Sscanf(idParam, "%d", &userID)
	return userID, err
}

// Simple in-memory store for users
var users = []User{}

func findUserByID(id int) (User, bool) {
	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}
	return User{}, false
}
