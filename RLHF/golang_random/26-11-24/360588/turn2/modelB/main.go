package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	// Replace this with your actual database or a secure storage mechanism
	authorizedUsers = map[string]string{
		"user1": "password1",
		"user2": "password2",
	}
)

func main() {
	r := gin.Default()

	r.POST("/auth/token", basicAuth)

	r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}

func basicAuth(c *gin.Context) {
	var u user
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate the user credentials
	if storedPassword, ok := authorizedUsers[u.Username]; ok {
		if storedPassword == u.Password {
			// Generate a JWT token (simplified example for demonstration purposes)
			token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
			c.JSON(http.StatusOK, gin.H{"access_token": token})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
}
