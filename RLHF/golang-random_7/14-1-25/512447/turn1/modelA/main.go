package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sendMessages(w *sync.WaitGroup, messageChannel chan<- string, latency time.Duration) {
	defer w.Done()
	for i := 0; i < 1000; i++ {
		messageChannel <- fmt.Sprintf("Message %d", i)
		time.Sleep(latency)
	}
}

func receiveMessages(w *sync.WaitGroup, messageChannel <-chan string) {
	defer w.Done()
	var receivedMessages []string
	for message := range messageChannel {
		receivedMessages = append(receivedMessages, message)
	}
	fmt.Println("Received Messages:", receivedMessages)
}

func main() {
	// Initialize variables
	latency := time.Millisecond * 100 // 100 milliseconds of latency
	messageChannel := make(chan string)
	waitGroup := &sync.WaitGroup{}

	// Start Goroutines
	waitGroup.Add(1)
	go sendMessages(waitGroup, messageChannel, latency)

	waitGroup.Add(1)
	go receiveMessages(waitGroup, messageChannel)

	// Wait for all Goroutines to finish
	waitGroup.Wait()

	// Analyze memory usage
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Memory Allocation:\n")
	fmt.Printf("Alloc: %d bytes\n", memStats.Alloc)
	fmt.Printf("TotalAlloc: %d bytes\n", memStats.TotalAlloc)
	fmt.Printf("Sys: %d bytes\n", memStats.Sys)
	fmt.Printf("Mallocs: %d\n", memStats.Mallocs)
	fmt.Printf("Frees: %d\n", memStats.Frees)
	fmt.Printf("HeapAlloc: %d bytes\n", memStats.HeapAlloc)
	fmt.Printf("HeapSys: %d bytes\n", memStats.HeapSys)
	fmt.Printf("HeapIdle: %d bytes\n", memStats.HeapIdle)
	fmt.Printf("HeapInuse: %d bytes\n", memStats.HeapInuse)
	fmt.Printf("HeapReleased: %d bytes\n", memStats.HeapReleased)
	fmt.Printf("HeapObjects: %d\n", memStats.HeapObjects)

	fmt.Println("Program completed.")
}
