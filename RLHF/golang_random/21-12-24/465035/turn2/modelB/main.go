package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func downloadFile(wg *sync.WaitGroup, fileURL string, fileName string) {
	defer wg.Done()
	response, err := http.Get(fileURL)
	if err != nil {
		fmt.Printf("Error downloading %s: %v\n", fileURL, err)
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body for %s: %v\n", fileURL, err)
		return
	}
	// You can further process the downloaded data here
	// For simplicity, just printing the file name as downloaded
	fmt.Printf("Downloaded: %s\n", fileName)
}

func main() {
	var wg sync.WaitGroup
	fileURLs := []string{
		"https://example.com/file1.txt",
		"https://example.com/file2.txt",
		"https://example.com/file3.txt",
		// Add more file URLs here for concurrent downloading
	}

	// Add a counter for each goroutine to be launched
	wg.Add(len(fileURLs))
	for _, fileURL := range fileURLs {
		// Extract the file name from the URL (you might need better file naming logic)
		fileName := fileURL[len(fileURL)-5:]
		go downloadFile(&wg, fileURL, fileName)
	}

	// The main goroutine will wait here for all worker goroutines to finish
	fmt.Println("Main thread waiting for all downloads to complete...")
	wg.Wait()
	fmt.Println("All downloads completed. Main thread exiting.")
}
