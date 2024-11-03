package main

import (
	"fmt"
	"time"
)

func main() {
	go receiveDataFromCollector()
	for {
		// Simulate analysis work
		time.Sleep(time.Second * 2)
	}
}

func receiveDataFromCollector() {
	for data := range analyzerChannel {
		fmt.Println("Analyzer: Received data:", data)
		// Perform data analysis here
	}
}
