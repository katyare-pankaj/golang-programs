package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// DataFetcher fetches data from an external source
type DataFetcher struct{}

func (df *DataFetcher) FetchData() int {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	return rand.Intn(100)
}

// PageRenderer renders a page using fetched data
type PageRenderer struct {
	cache map[int]string
	mu    sync.Mutex
}

func (pr *PageRenderer) RenderPage(dataFetcher *DataFetcher) string {
	data := dataFetcher.FetchData()

	// Check if the page is in the cache
	pr.mu.Lock()
	cachedPage, ok := pr.cache[data]
	pr.mu.Unlock()

	if ok {
		fmt.Println("Page found in cache.")
		return cachedPage
	}

	// Render the page and cache it
	page := fmt.Sprintf("Rendered page with data: %d", data)
	pr.mu.Lock()
	pr.cache[data] = page
	pr.mu.Unlock()

	fmt.Println("Page rendered and cached.")
	return page
}

func main() {
	dataFetcher := &DataFetcher{}
	pageRenderer := &PageRenderer{cache: make(map[int]string)}

	// Simulate multiple requests
	for i := 0; i < 10; i++ {
		go func() {
			page := pageRenderer.RenderPage(dataFetcher)
			fmt.Println("Rendered Page:", page)
		}()
	}

	// Wait for all goroutines to complete
	time.Sleep(2 * time.Second)
}
