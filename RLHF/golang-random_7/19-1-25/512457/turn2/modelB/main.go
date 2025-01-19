package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	r := gin.Default()
	r.POST("/users", createUser)
	r.Run(":8080")
}

func createUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	// Add your user creation logic here
	fmt.Println("User created:", req.Username, req.Email)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func handleError(c *gin.Context, statusCode int, err error) {
	apiErr, ok := err.(ApiError)
	if !ok {
		apiErr = ApiError{
			Code:    statusCode,
			Message: err.Error(),
		}
	}

	c.JSON(statusCode, gin.H{
		"error": apiErr,
	})
}
