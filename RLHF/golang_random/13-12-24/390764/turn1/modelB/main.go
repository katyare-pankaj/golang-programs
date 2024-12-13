package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	wg   sync.WaitGroup
	urls = []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		"https://golang.org",
	}
)

func fetch(url string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body for %s: %v\n", url, err)
		return
	}

	fmt.Printf("Fetched %s: %d bytes\n", url, len(body))
}

func main() {
	wg.Add(len(urls))

	for _, url := range urls {
		go fetch(url) // Start a Goroutine for each URL
	}

	wg.Wait() // Wait for all Goroutines to finish

	fmt.Println("All URLs fetched.")
}
