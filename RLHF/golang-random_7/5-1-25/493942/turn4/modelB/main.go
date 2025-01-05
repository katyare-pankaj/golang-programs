package main

import (
	"fmt"
	"net/http"
)

func logRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method: %s, URL: %s\n", r.Method, r.URL.Path)
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logRequest(w, r)
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe(":8080", nil)
}
