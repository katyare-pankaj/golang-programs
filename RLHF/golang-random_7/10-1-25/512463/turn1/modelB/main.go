package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func makeAPICall(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		// Handle error gracefully
		log.Printf("Error making API call to %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		// Handle error gracefully
		log.Printf("API call to %s returned status code %d", url, resp.StatusCode)
		return
	}
	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Handle error gracefully
		log.Printf("Error reading response body from %s: %v", url, err)
		return
	}

	// Print the response body for successful calls
	fmt.Printf("Response from %s: %s\n", url, string(body))
}

func main() {
	urls := []string{"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/4",
		"https://jsonplaceholder.typicode.com/posts/5"}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	// Perform API calls concurrently
	for _, url := range urls {
		go makeAPICall(url, &wg)
	}

	wg.Wait()
	// All API calls have completed
	fmt.Println("All API calls completed.")
}
