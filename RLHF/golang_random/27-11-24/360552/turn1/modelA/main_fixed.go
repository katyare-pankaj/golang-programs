package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// proxy handles incoming requests and proxies them to the specified backend URL.
func proxy(backend string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the backend URL
		backendURL, err := url.Parse(backend)
		if err != nil {
			http.Error(w, "Invalid backend URL", http.StatusInternalServerError)
			return
		}

		// Build the proxied request
		proxyURL := backendURL.ResolveReference(r.URL)
		req, err := http.NewRequestWithContext(r.Context(), r.Method, proxyURL.String(), r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating request: %v", err), http.StatusInternalServerError)
			return
		}

		// Copy headers
		for k, v := range r.Header {
			req.Header[k] = v
		}
		req.Header.Set("Host", backendURL.Host)

		// Make the request to the backend
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error making request: %v", err), http.StatusServiceUnavailable)
			return
		}
		defer res.Body.Close()

		// Copy the response back to the client
		for k, v := range res.Header {
			w.Header()[k] = v
		}
		w.WriteHeader(res.StatusCode)
		_, err = io.Copy(w, res.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error copying response: %v", err), http.StatusInternalServerError)
		}
	}
}

func main() {
	backend := "http://localhost:8080" // Replace with actual backend URL

	log.Println("Starting API Gateway...")
	http.HandleFunc("/api/v1/", proxy(backend))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
