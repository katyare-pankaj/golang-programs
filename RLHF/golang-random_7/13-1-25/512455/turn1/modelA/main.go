package main

import (
	"fmt"
	"net/http"
	"sync"
)

func getData(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter after this function returns
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting data from %s: %v\n", url, err)
		return
	}
	defer response.Body.Close() // Close the response body

	body, err := http.GetString(url)
	if err != nil {
		fmt.Printf("Error reading body from %s: %v\n", url, err)
		return
	}

	// Do something with the data
	fmt.Printf("Data from %s:\n%s\n", url, body)
}

func main() {
	var wg sync.WaitGroup

	// Example travel APIs URLs
	apis := []string{
		"http://api.travelservice1.com/data",
		"http://api.travelservice2.com/data",
		"http://api.travelservice3.com/data",
	}

	// Initialize the WaitGroup
	wg.Add(len(apis))

	// Launch Goroutines to fetch data from each API
	for _, url := range apis {
		go getData(url, &wg)
	}

	// Wait for all Goroutines to complete
	wg.Wait()

	fmt.Println("All tasks have completed.")
}
