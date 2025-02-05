package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	port         = ":8080"
	numRequests  = 100
	concurrency  = 10
	maxLatencyMs = 100 // Maximum simulated latency in milliseconds
)

func main() {
	// Step 1: Set up the HTTP server
	http.HandleFunc("/benchmark", benchmarkHandler)
	go func() {
		log.Printf("Starting server at %s\n", port)
		if err := http.ListenAndServe(port, nil); err != nil {
			log.Fatalf("Could not start server: %s\n", err)
		}
	}()

	// Give server a second to start
	time.Sleep(1 * time.Second)

	// Step 2: Send requests
	var wg sync.WaitGroup
	wg.Add(numRequests)
	startTime := time.Now()

	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			sendRequest()
		}()
	}

	wg.Wait()
	endTime := time.Now()

	// Step 3: Output the total time taken
	fmt.Printf("Total time for %d requests with concurrency %d: %v\n", numRequests, concurrency, endTime.Sub(startTime))
}

func benchmarkHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate network latency
	latency := time.Duration(rand.Intn(maxLatencyMs)) * time.Millisecond
	time.Sleep(latency)

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Response from server"))
}

func sendRequest() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost"+port+"/benchmark", nil)
	if err != nil {
		log.Fatalf("Could not create request: %s\n", err)
	}

	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Request failed: %s\n", err)
		return
	}
	defer resp.Body.Close()

	duration := time.Since(start)
	log.Printf("Request took %s\n", duration)
}
