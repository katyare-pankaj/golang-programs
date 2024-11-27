package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Middleware for basic authentication
	router.Use(BasicAuth())

	// Define routes
	router.GET("/api/v1/users", GetUsers)
	router.POST("/api/v1/users", CreateUser)

	// Start the server on port 8000
	err := router.Run(":8000")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// BasicAuth Middleware
func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Basic Authentication credentials
		username, password, ok := c.Request.BasicAuth()
		if !ok || username != "your_username" || password != "your_password" {
			// Request Basic Authentication if credentials are not provided
			c.Header("WWW-Authenticate", "Basic realm=\"Restricted\"")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Valid credentials, allow the request to proceed
		c.Next()
	}
}

// User struct for demonstration purposes
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Doe"},
}

// GetUsers Handler
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// CreateUser Handler
func CreateUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Add the new user to the users slice
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}
