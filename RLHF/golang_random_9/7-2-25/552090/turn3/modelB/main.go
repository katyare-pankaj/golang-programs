package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func downloadFile(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simple file download logic using HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading file from %s: %s\n", url, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body from %s: %s\n", url, err)
		return
	}
	// You can do something with the downloaded file content,
	// like saving it to a local file, here just print the length
	fmt.Printf("Downloaded %d bytes from %s\n", len(body), url)
}

func main() {
	urls := []string{
		"https://example.com/file1.txt",
		"https://example.com/file2.txt",
		"https://example.com/file3.txt",
		"https://example.com/file4.txt",
		// Add more URLs as needed
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	// Launch goroutines for each download
	for _, url := range urls {
		go downloadFile(url, &wg)
	}

	wg.Wait() // Wait for all downloads to complete
	fmt.Println("All downloads completed.")
}
