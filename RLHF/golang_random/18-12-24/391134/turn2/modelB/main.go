package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

const (
	numAPICalls   = 10
	apiTimeout    = 2 * time.Second
	totalDuration = 10 * time.Second
)

func fetchData(ctx context.Context, url string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			log.Printf("Request to %s cancelled: %v", url, ctx.Err())
		default:
			log.Printf("Error fetching data from %s: %v", url, err)
		}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code from %s: %d", url, resp.StatusCode)
		return
	}

	// Simulate working with the response body
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Printf("Successfully fetched data from: %s\n", url)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), totalDuration)
	defer cancel()

	wg := &sync.WaitGroup{}

	// Start API calls concurrently
	for i := 0; i < numAPICalls; i++ {
		url := fmt.Sprintf("https://example.com/api/%d", i)
		wg.Add(1)
		go fetchData(ctx, url, wg)
	}

	// Wait for all API calls to complete or the context to be cancelled
	wg.Wait()

	select {
	case <-ctx.Done():
		log.Printf("Execution timed out: %v", ctx.Err())
	default:
		log.Println("All API calls completed successfully.")
	}
}
