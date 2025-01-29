package main

import (
	"fmt"
	"net/http"
	"sync"
)

// processURL fetches a URL and ensures cleanup of any acquired resources.
func processURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	// Immediate defer ensures the resource is always cleaned up.
	defer func() {
		if resp != nil {
			resp.Body.Close()
			fmt.Println("Response body closed for URL:", url)
		}
	}()

	if err != nil {
		fmt.Printf("Error fetching URL %s: %v\n", url, err)
		return
	}

	// Process the response
	fmt.Printf("Successfully fetched the URL %s, status code: %d\n", url, resp.StatusCode)
	// Other processing logic...
}

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"http://example.com",
		"http://golang.org",
		"http://invalid.fakeurl",
	}

	for _, url := range urls {
		wg.Add(1)
		go processURL(url, &wg)
	}

	wg.Wait()
	fmt.Println("All URL processing completed!")
}
