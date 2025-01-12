package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func fetchData(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading body from %s: %v\n", url, err)
		return
	}

	fmt.Printf("Fetched %s with %d bytes of data\n", url, len(data))
}

func main() {
	var wg sync.WaitGroup

	// List of URLs to make API calls to
	urls := []string{
		"https://api.example.com/data1",
		"https://api.example.com/data2",
		"https://api.example.com/data3",
		"https://api.example.com/data4",
	}

	// Add one to the WaitGroup for each API call
	wg.Add(len(urls))

	// Make the API calls concurrently
	for _, url := range urls {
		go fetchData(url, &wg)
	}

	// Wait for all API calls to complete
	wg.Wait()

	fmt.Println("All API calls completed")
}
