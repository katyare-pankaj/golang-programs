package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	numWorkers = 4   // Number of goroutines for concurrent processing
	dataBuffer = 100 // Buffer size for the channel
)

func processData(data chan string) {
	for item := range data {
		// Simulate data processing time
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("Processed:", item)
	}
}

func main() {
	dataChan := make(chan string, dataBuffer)

	// Start worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			processData(dataChan)
			wg.Done()
		}()
	}

	// Sample incoming data from HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dataItem := r.URL.Query().Get("data")
		if dataItem != "" {
			dataChan <- dataItem
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
	})

	fmt.Println("Server started. Listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

	// Close the channel to signal workers to finish
	close(dataChan)
	wg.Wait()
}
