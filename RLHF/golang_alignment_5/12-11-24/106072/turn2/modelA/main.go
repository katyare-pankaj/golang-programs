package main

import (
	"fmt"
	"time"
)

const (
	apiRequestLimitPerSecond = 5 // Set the desired API request limit per second
)

var (
	requestCount     int
	lastRequestTime  time.Time
	throttleDuration time.Duration
)

func init() {
	throttleDuration = time.Second / time.Duration(apiRequestLimitPerSecond)
}

func allowRequest() bool {
	currentTime := time.Now()
	if currentTime.Sub(lastRequestTime) >= throttleDuration {
		requestCount = 0
		lastRequestTime = currentTime
	}

	if requestCount < apiRequestLimitPerSecond {
		requestCount++
		return true
	}

	return false
}

func main() {
	for i := 0; i < 20; i++ {
		if allowRequest() {
			fmt.Println("Making API request...")
			// Simulate API call duration
			time.Sleep(200 * time.Millisecond)
		} else {
			fmt.Println("Rate limit exceeded. Waiting...")
			time.Sleep(throttleDuration)
		}
	}
}
