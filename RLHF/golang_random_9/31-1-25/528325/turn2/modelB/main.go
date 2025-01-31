package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const (
	numRequests = 1000000 // Total number of requests to process
	numWorkers  = 100     // Number of goroutines to handle requests
	requestSize = 1024    // Size of each request payload in bytes
)

var (
	wg            sync.WaitGroup
	startTime     time.Time
	totalRequests int64
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	defer wg.Done()
	// Simulate request processing
	time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
	totalRequests++
}
func worker(id int) {
	for {
		// Create a new request for processing
		req, err := http.NewRequest("GET", "http://localhost:8080", nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}
		// Simulate request payload size
		req.Body = http.NoBody // You can use ioutil.NopCloser(bytes.NewReader(make([]byte, requestSize))) for payload
		// Process the request
		wg.Add(1)
		http.DefaultClient.Do(req)
	}
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Start the HTTP server
	go func() {
		http.HandleFunc("/", requestHandler)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Error starting HTTP server:", err)
		}
	}()
	// Wait for server to start
	time.Sleep(time.Second)
	startTime = time.Now()
	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go worker(i)
	}
	// Wait for all requests to be processed
	wg.Wait()
	elapsedTime := time.Since(startTime)
	fmt.Println("Total requests processed:", totalRequests)
	fmt.Println("Elapsed time:", elapsedTime)
	fmt.Println("Throughput:", float64(totalRequests)/elapsedTime.Seconds(), "requests/sec")
	// Get GC statistics
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Println("GC Cycles:", stats.NumGC)
	fmt.Println("GC Pause Total:", stats.PauseTotalNs/1e6, "ms")
}
