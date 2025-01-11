package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Result struct {
	URL        string
	StatusCode int
	Error      error
	Body       string
}

func makeAPICall(url string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	result := Result{URL: url}

	if err != nil {
		result.Error = err
	} else {
		defer resp.Body.Close()
		result.StatusCode = resp.StatusCode
		if resp.StatusCode == http.StatusOK {
			body, _ := ioutil.ReadAll(resp.Body)
			result.Body = string(body)
		}
	}
	results <- result
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/4",
		"https://jsonplaceholder.typicode.com/posts/5",
		"https://jsonplaceholder.typicode.com/posts/6", // Added a URL that will fail
	}

	results := make(chan Result, len(urls))
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go makeAPICall(url, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var successfulResults []Result
	var failedResults []Result

	for result := range results {
		if result.Error != nil {
			failedResults = append(failedResults, result)
		} else {
			successfulResults = append(successfulResults, result)
		}
	}

	// Process successful results here
	fmt.Println("Successful Results:")
	for _, result := range successfulResults {
		fmt.Printf("URL: %s, Status Code: %d, Body: %s\n", result.URL, result.StatusCode, result.Body)
	}

	// Process failed results here
	fmt.Println("\nFailed Results:")
	for _, result := range failedResults {
		fmt.Printf("URL: %s, Error: %v\n", result.URL, result.Error)
	}
}
