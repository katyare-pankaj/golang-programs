package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Middleware function to log requests.
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Increment the WaitGroup to account for the goroutine.
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			// Perform some work...
			time.Sleep(time.Second)
			defer wg.Done()
		}()
		// Execute the next handler in the chain.
		next.ServeHTTP(w, r)
		// Wait for the goroutine to complete before returning.
		wg.Wait()
	})
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	http.ListenAndServe(":8080", loggerMiddleware(http.DefaultServeMux))
}
