package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User struct represents the user entity
type User struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func createUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user data",
		})
		return
	}

	// Add additional user creation logic here

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func main() {
	router := gin.Default()
	router.POST("/users", createUser)
	router.Run(":8080")
}
