package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			next.ServeHTTP(w, r)
		}()
		wg.Wait()
		duration := time.Since(start)
		fmt.Printf("%s\t%s\t%v\n", r.Method, r.URL.Path, duration)
	})
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second) // Simulate a slow request
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Slow handler responded!")
}

func fastHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Fast handler responded!")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/slow", loggingMiddleware(http.HandlerFunc(slowHandler)))
	mux.Handle("/fast", loggingMiddleware(http.HandlerFunc(fastHandler)))
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", mux)
}
