package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define your JSON response struct
type Response struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Name: "Alice",
			Age:  25,
		}
		// Marshal the response struct to JSON
		data, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")
		// Write the JSON response to the client
		w.Write(data)
	})
	// Start the HTTP server
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
