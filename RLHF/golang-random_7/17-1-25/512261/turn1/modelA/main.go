package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

// A third-party library function that doesn't handle defer properly.
func thirdPartyFunction(url string, w io.Writer) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(w, res.Body)
	return err
}

// A wrapper function that ensures resources are properly closed.
func safeThirdPartyFunction(url string, w io.Writer) error {
	var err error
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err2 := res.Body.Close(); err2 != nil {
			log.Printf("Error closing response body: %v", err2)
		}
	}()

	_, err = io.Copy(w, res.Body)
	return err
}

// Main function to demonstrate concurrent usage.
func main() {
	var wg sync.WaitGroup
	numGoroutines := 100
	url := "https://example.com"

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: Starting with unsafe third-party function...\n", id)
			start := time.Now()
			err := thirdPartyFunction(url, io.Discard)
			if err != nil {
				log.Printf("Goroutine %d: Error: %v\n", id, err)
			}
			duration := time.Since(start)
			fmt.Printf("Goroutine %d: Completed with unsafe third-party function in %s\n", id, duration)
		}(i)
	}

	fmt.Println("Starting safe third-party function calls...")
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: Starting with safe third-party function...\n", id)
			start := time.Now()
			err := safeThirdPartyFunction(url, io.Discard)
			if err != nil {
				log.Printf("Goroutine %d: Error: %v\n", id, err)
			}
			duration := time.Since(start)
			fmt.Printf("Goroutine %d: Completed with safe third-party function in %s\n", id, duration)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed.")
}
