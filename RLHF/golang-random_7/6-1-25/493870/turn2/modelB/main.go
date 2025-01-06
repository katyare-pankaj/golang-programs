package main

import (
	"fmt"
	"net/http"

	lru "github.com/hashicorp/golang-lru"
)

func main() {
	// Create a new cache with a capacity of 100 status codes
	cache, _ := lru.New(100)

	// Load HTTP status codes and their descriptions into the cache
	for statusCode := http.StatusContinue; statusCode <= http.StatusNetworkAuthenticationRequired; statusCode++ {
		statusText := http.StatusText(statusCode)
		if len(statusText) > 0 {
			cache.Add(statusCode, statusText)
		}
	}

	// Display the formatted list of HTTP status codes and their descriptions
	fmt.Println("HTTP Status Codes:")
	cache.ForEach(func(key, value interface{}) bool {
		statusCode := key.(int)
		statusText := value.(string)
		fmt.Printf("%-3d: %s\n", statusCode, statusText)
		return true
	})
}
