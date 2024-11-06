package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// SampleData represents a sample data point for visualization
type SampleData struct {
	Name      string `json:"name"`
	Value     int    `json:"value"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/data", dataHandler).Methods("GET")

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Generate sample data (replace this with actual data processing)
	data := generateSampleData(10000)

	// Optimize: Perform data processing in a Web Worker (not directly in Go for web apps)
	//start := time.Now()
	//processedData := optimizeDataProcessing(data)
	//elapsed := time.Since(start)
	//fmt.Printf("Data processing time: %s\n", elapsed)

	// Respond with the data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func generateSampleData(size int) []SampleData {
	data := make([]SampleData, 0, size)
	for i := 0; i < size; i++ {
		data = append(data, SampleData{
			Name:      fmt.Sprintf("Item %d", i),
			Value:     i * 10,
			Timestamp: time.Now().UnixNano(),
		})
	}
	return data
}

// **Note:** The 'optimizeDataProcessing' function is not implemented here as it would require using Web Workers,
// which is not directly supported in Go for web applications without additional libraries.
