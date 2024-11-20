package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

// Database Stream Response
type DatabaseStreamResponse struct {
	ID        int       `json:"id"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

// Define rate limit structure
type rateLimit struct {
	limit     int
	remaining int
	reset     time.Time
}

// Database stream function (Simulates streaming data from a NoSQL database)
func databaseStream(ch chan<- DatabaseStreamResponse) {
	for i := 1; i <= 100; i++ {
		data := fmt.Sprintf("Sample Data %d", i)
		response := DatabaseStreamResponse{ID: i, Data: data, CreatedAt: time.Now()}
		ch <- response // Send the data through the channel
	}
	close(ch)
}

// Rate limit middleware
func rateLimiter(c *gin.Context, rateLimiterCache *cache.Cache) {
	// Get the IP address of the client
	clientIP := c.ClientIP()

	// Check if the client has a rate limit in the cache
	rateLimitInterface, found := rateLimiterCache.Get(clientIP)
	if !found {
		// If the client does not have a rate limit, set a new one
		rateLimiterCache.Set(clientIP, rateLimit{limit: 10, remaining: 10, reset: time.Now().Add(time.Second)}, cache.DefaultExpiration)
	}

	rateLimit := rateLimitInterface.(rateLimit)

	// Check if the client has exceeded the rate limit
	if rateLimit.remaining <= 0 {
		c.JSON(429, gin.H{
			"error": "Too Many Requests",
		})
		c.Abort()
		return
	}

	// Decrement the remaining requests
	rateLimit.remaining--

	// Update the rate limit in the cache
	rateLimiterCache.Set(clientIP, rateLimit, cache.DefaultExpiration)

	c.Next()
}

func main() {
	router := gin.Default()

	// Initialize rate limiter cache
	rateLimiterCache := cache.New(5*time.Minute, 10*time.Minute)

	// Use the rate limiter middleware for all API endpoints
	router.Use(func(c *gin.Context) {
		rateLimiter(c, rateLimiterCache)
	})

	router.GET("/api/stream", func(c *gin.Context) {
		// Start the database stream
		ch := make(chan DatabaseStreamResponse)
		go databaseStream(ch)

		c.Stream(func(w gin.ResponseWriter) bool {
			for response := range ch {
				// Format the response as JSON
				if err := c.JSON(200, response); err != nil {
					return false // Stream error, return false
				}

				// Flush the writer to send the response immediately
				c.Writer.Flush()

				// Add a small delay to simulate real-time data streaming
				time.Sleep(100 * time.Millisecond)
			}

			return true // Stream completed, return true
		})
	})

	router.Run(":8080")
}
