package main

import (
	"log"
	"net/http"
	"net/http/gzip"
)

func main() {
	http.HandleFunc("/", gzippedHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func gzippedHandler(w http.ResponseWriter, r *http.Request) {
	// Your actual handler code here
	// ...
	// Compress the response automatically
	gzipHandler := gzip.NewHandler(w)
	gzipHandler.ServeHTTP(w, r)
}
