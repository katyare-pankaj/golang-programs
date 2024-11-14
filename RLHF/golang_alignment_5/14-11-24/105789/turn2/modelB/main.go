package main

import (
	"time"

	"github.com/micro/go-micro/client"
)

func main() {
	// Create a client
	c := client.NewClient()

	// Define a retry policy
	retry := client.Retries(3, 100*time.Millisecond)

	// Set the retry policy on the client
	c.Options().SetRetry(retry)

	// Call the service with retries
	err := callService(c)
	if err != nil {
		panic(err)
	}
}
