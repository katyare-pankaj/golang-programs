package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Result struct {
	URL      string
	Response string
	Error    error
}

func makeAPICall(url string, wg *sync.WaitGroup, resultChannel chan<- Result) {
	defer wg.Done()

	var result Result
	resp, err := http.Get(url)
	if err != nil {
		// Handle error and set Result.Error
		log.Printf("Error making API call to %s: %v", url, err)
		result.URL = url
		result.Error = err
	} else {
		defer resp.Body.Close()

		// Check if the response status code is 200 OK
		if resp.StatusCode != http.StatusOK {
			// Handle error and set Result.Error
			log.Printf("API call to %s returned status code %d", url, resp.StatusCode)
			result.URL = url
			result.Error = fmt.Errorf("status code %d", resp.StatusCode)
		} else {
			// Read the response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// Handle error and set Result.Error
				log.Printf("Error reading response body from %s: %v", url, err)
				result.URL = url
				result.Error = err
			} else {
				// Store successful response
				result.URL = url
				result.Response = string(body)
			}
		}
	}

	// Send the Result to the channel
	resultChannel <- result
}

func processResults(results []Result) {
	for _, result := range results {
		if result.Error != nil {
			log.Printf("Error for %s: %v", result.URL, result.Error)
		} else {
			// Process the successful response
			// For example, parse JSON
			var post Post
			if err := json.Unmarshal([]byte(result.Response), &post); err != nil {
				log.Printf("Error parsing JSON for %s: %v", result.URL, err)
			} else {
				fmt.Printf("Successful response from %s: %+v\n", result.URL, post)
			}
		}
	}
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/4",
		"https://jsonplaceholder.typicode.com/posts/5",
	}

	resultChannel := make(chan Result, len(urls))
	var wg sync.WaitGroup
	wg.Add(len(urls))

	// Perform API calls concurrently
	for _, url := range urls {
		go makeAPICall(url, &wg, resultChannel)
	}

	// Collect results from the channel
	var results []Result
	for i := 0; i < len(urls); i++ {
		results = append(results, <-resultChannel)
	}

	wg.Wait()
	// All API calls have completed
	fmt.Println("All API calls completed.")

	// Process successful and failed results separately
	processResults(results)
}
