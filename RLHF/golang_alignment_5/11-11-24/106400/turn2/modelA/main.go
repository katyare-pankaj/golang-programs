// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

// HelloHandler handles HTTP requests to /hello
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	message := fmt.Sprintf("Hello, %s!", name)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, message)
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
