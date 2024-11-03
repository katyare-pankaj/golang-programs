package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		data := rand.Intn(100)
		fmt.Println("Sensor Data:", data)
		// Simulate sending data to the data analyzer microservice over a local network connection (e.g., Unix socket or in-memory channel)
		sendDataToAnalyzer(data)
		time.Sleep(time.Second * 1)
	}
}

func sendDataToAnalyzer(data int) {
	// In a real scenario, you would use IPC (Inter-Process Communication) here
	// For simplicity, we'll use a channel in this example
	analyzerChannel <- data
}

var analyzerChannel = make(chan int)
