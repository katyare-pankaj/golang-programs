package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// Database connection details
var db *sql.DB

const dbDSN = "root:password@tcp(localhost:3306)/mydatabase?parseTime=true"

// Redis client
var redisClient *redis.Client

func init() {
	// Initialize the database connection
	db, err := sql.Open("mysql", dbDSN)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	// Ping the database to test connection
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatalf("Failed to ping MySQL: %v", err)
	}
	fmt.Println("MySQL connection established.")

	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Redis password (if any)
		DB:       0,
	})

	// Ping Redis to test connection
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("Redis connection established.")
}

func healthCheck(c *gin.Context) {
	// Check MySQL availability
	err := db.Ping(context.Background())
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "MySQL service is down",
		})
		return
	}

	// Check Redis availability
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Redis service is down",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All services are healthy",
	})
}

func main() {
	r := gin.Default()

	// Register the health check endpoint
	r.GET("/health", healthCheck)

	// Start the server
	log.Fatal(r.Run(":8080"))
}
