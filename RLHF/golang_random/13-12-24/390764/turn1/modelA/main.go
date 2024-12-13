package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the function completes

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	fmt.Printf("Fetched %d bytes from %s\n", len(body), url)
}

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"http://example.com",
		"http://example.org",
		"http://example.net",
	}

	for _, url := range urls {
		wg.Add(1)             // Increment the WaitGroup counter
		go fetchURL(url, &wg) // Start a new Goroutine
	}

	wg.Wait() // Wait for all Goroutines to finish
	fmt.Println("All URLs fetched.")
}
