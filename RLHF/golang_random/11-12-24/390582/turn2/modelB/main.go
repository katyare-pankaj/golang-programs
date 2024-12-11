package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func slowHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printf("Handling request with deadline: %v", ctx.Deadline())

	select {
	case <-time.After(5 * time.Second):
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Response after 5 seconds"))
	case <-ctx.Done():
		log.Printf("Request cancelled: %v", ctx.Err())
		w.WriteHeader(http.StatusGatewayTimeout)
		w.Write([]byte("Request timed out"))
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set a 3-second timeout for the request
		ctx, cancel := context.WithDeadline(r.Context(), time.Now().Add(3*time.Second))
		defer cancel()
		r = r.WithContext(ctx)

		slowHandler(w, r)
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
