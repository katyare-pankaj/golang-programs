package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// proxy handles incoming requests and proxies them to the specified backend URL.
func proxy(backend string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/") {
			newPath := fmt.Sprintf("%s%s", backend, r.URL.Path)
			r.URL.Path = newPath
			r.URL.Scheme = "http"
			r.URL.Host = backend[strings.IndexByte(backend, ':')+1:]
		} else {
			r.URL.Path = fmt.Sprintf("%s/%s", backend, r.URL.Path)
		}
		r.URL.Host = backend
		client := &http.Client{}
		req, err := http.NewRequestWithContext(r.Context(), r.Method, r.URL.String(), r.Body)
		if err != nil {
			log.Fatalf("Error creating request: %v", err)
		}

		for k, v := range r.Header {
			req.Header[k] = v
		}
		req.Header.Set("Host", r.Host)

		res, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		defer res.Body.Close()

		for k, v := range res.Header {
			w.Header()[k] = v
		}
		w.WriteHeader(res.StatusCode)
		_, err = http.Ioutil.Copy(w, res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
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
