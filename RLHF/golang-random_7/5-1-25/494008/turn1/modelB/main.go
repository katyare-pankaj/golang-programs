package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Response struct {
	Body  []byte
	Error error
}

// Function to fetch data
func fetch(url string, c chan<- Response) {
	resp, err := http.Get(url)
	if err != nil {
		c <- Response{Error: err}
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c <- Response{Error: err}
		return
	}
	c <- Response{Body: body}
}

// Main function
func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.bbc.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
		"https://www.wikipedia.org",
	}

	results := make(chan Response, len(urls))
	//Scrape the data from urls using goroutines
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fetch(url, results)
		}(url)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(results)
	// Prints scraped data from each site
	for result := range results {
		if result.Error != nil {
			fmt.Println("Error:", result.Error)
		} else {
			fmt.Println(string(result.Body))
		}
	}
}
