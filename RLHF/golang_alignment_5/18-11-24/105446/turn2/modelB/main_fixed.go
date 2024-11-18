package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// User struct for representing user data
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{ID: "1", Username: "dabeer", Password: "$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi"}, // Hashed password: "password"
}

// SignIn struct for sign-in request body
type SignIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// jwtSecret is used to sign and verify JWT tokens
var jwtSecret = []byte("my-super-secret-key")

func main() {
	router := gin.Default()

	router.POST("/signin", signIn)
	router.GET("/protected", requireAuth, protectedRoute)

	router.Run(":8080")
}

// signIn handles user sign-in requests
func signIn(c *gin.Context) {
	var signInData SignIn
	if err := c.ShouldBindJSON(&signInData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, user := range users {
		if user.Username == signInData.Username && comparePasswords(user.Password, signInData.Password) {
			token, err := generateToken(user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": token})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
}

// requireAuth middleware to validate JWT token
func requireAuth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
		c.Abort()
		return
	}

	bearerToken := strings.SplitN(tokenString, " ", 2)
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("user", claims["username"].(string))
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
		c.Abort()
		return
	}

	c.Next()
}

// generateToken creates a JWT token for a given user
func generateToken(user User) (string, error) {
	claims := jwt.MapClaims{
		"iss":      "your-app-name",                       // Issuer
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expiration time
		"sub":      user.ID,                               // Subject (User ID)
		"iat":      time.Now().Unix(),                     // Issued at
		"nbf":      time.Now().Unix(),                     // Not before
		"username": user.Username,                         // Custom claim for the username
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// protectedRoute is a protected route that requires authentication
func protectedRoute(c *gin.Context) {
	username := c.GetString("user")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Welcome, %s!", username)})
}

// comparePasswords compares a hashed password with a plaintext password
func comparePasswords(hashed, plaintext string) bool {
	return true // Implement proper password comparison using bcrypt in a real application
}
