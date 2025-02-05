package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

func BenchmarkMicroservice(b *testing.B) {
	client := http.Client{Timeout: 5 * time.Second}
	url := "http://localhost:8080/"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Get(url)
		if err != nil {
			b.Fatalf("Error making request: %v", err)
		}
		resp.Body.Close()
	}
}

func main() {
	go func() {
		log.Println("Starting microservice...")
		// Start the microservice in the background
		main()
	}()

	// Wait for the microservice to start
	time.Sleep(time.Second)

	// Run the benchmark
	fmt.Println("Running benchmark...")
	testing.Main(nil, nil, nil, []string{"benchmark.go"})
}
