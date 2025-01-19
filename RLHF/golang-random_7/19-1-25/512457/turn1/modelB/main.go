package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// User represents a user resource in the API
type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"gte=0"`
	Location string `json:"location"`
}

// VersionedAPIResponse is a wrapper for API responses to include version information
type VersionedAPIResponse struct {
	APIVersion string      `json:"api_version"`
	Data       interface{} `json:"data"`
	Error      *APIError   `json:"error,omitempty"`
}

// APIError represents an error in the API response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func main() {
	r := gin.Default()

	// Register API routes with versioning
	v1 := r.Group("/v1")
	{
		v1.POST("/users", createUser)
	}

	r.Run(":8080") // Listen and serve on port 8080
}

func createUser(c *gin.Context) {
	var user User

	// Bind the JSON request body to the User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Validate the user data
	if err := validate.Struct(user); err != nil {
		handleError(c, http.StatusBadRequest, "Validation error", err)
		return
	}

	// Perform business logic here
	// ...

	// Return the created user with version information
	c.JSON(http.StatusCreated, VersionedAPIResponse{
		APIVersion: "v1",
		Data:       user,
	})
}

func handleError(c *gin.Context, statusCode int, message string, err error) {
	// Log the error (optional)
	fmt.Println(err)

	// Create an APIError response with the appropriate error code and message
	apiError := APIError{
		Code:    statusCode,
		Message: message,
	}

	// Return the error response with version information
	c.JSON(statusCode, VersionedAPIResponse{
		APIVersion: "v1",
		Error:      &apiError,
	})
}
