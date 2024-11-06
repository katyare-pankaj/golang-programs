// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize router and routes
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/data.json", dataHandler)

	// Start the server
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the HTML file containing the VEGA-Lite visualization
	http.ServeFile(w, r, "index.html")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the JSON data for the visualization
	http.ServeFile(w, r, "data.json")
}
