package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Handle a single HTTP request concurrently
func handleHTTPRequest(w http.ResponseWriter, r *http.Request, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 500) // Simulate processing time
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go handleHTTPRequest(w, r, wg)
		wg.Wait()
	})

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
