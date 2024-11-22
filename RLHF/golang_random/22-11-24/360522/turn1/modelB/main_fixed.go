package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

type limiter struct {
	rate  int // maximum number of requests per second
	burst int // maximum number of requests burst
	cache *cache.Cache
	mux   sync.Mutex
}

func newLimiter(rate, burst int) *limiter {
	return &limiter{
		rate:  rate,
		burst: burst,
		cache: cache.New(1*time.Minute, 10*time.Minute), // Set the cache expiration time as needed
	}
}
func (l *limiter) Allow(key string) bool {
	l.mux.Lock()
	defer l.mux.Unlock()

	// Get the current request count for the key
	count, found := l.cache.Get(key)
	if !found {
		// Initialize the request count for the key
		count = 0
		l.cache.Set(key, count, cache.NoExpiration)
	}

	// Type assertion to int
	countInt, ok := count.(int)
	if !ok {
		// Handle the case where the count is not of type int
		return false
	}

	// Calculate the maximum number of requests allowed in the current second
	max := l.rate + l.burst

	// Check if the current request count is less than the maximum allowed
	if countInt < max {
		// Increment the request count and update the cache
		l.cache.Set(key, countInt+1, cache.NoExpiration)
		return true
	}

	return false
}

func main() {
	r := gin.Default()

	// Initialize the rate limiter with the desired rate and burst
	limiter := newLimiter(5, 10) // 5 requests per second with a burst of 10 requests

	// Define the API endpoint with the rate limiter middleware
	r.GET("/api/data", func(c *gin.Context) {
		// Get the user ID from the request context (or use any other unique identifier for the user)
		userID := c.GetString("user_id")

		// Check if the user is allowed to make a request
		if !limiter.Allow(userID) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Rate limit exceeded",
			})
			return
		}

		// Handle the API request here
		c.JSON(http.StatusOK, gin.H{
			"message": "Success",
		})
	})

	log.Fatal(r.Run(":8080"))
}
