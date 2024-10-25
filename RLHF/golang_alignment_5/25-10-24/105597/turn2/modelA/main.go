package main

import (
	"fmt"
	"net/http"
	"sort"
)

// Using a slice for in-memory data storage
var inMemoryData []int

func init() {
	// Populate inMemoryData with some sample data for sorting example
	for i := 0; i < 100000; i++ {
		inMemoryData = append(inMemoryData, i)
	}
}

func handleSortData(w http.ResponseWriter, r *http.Request) {
	// Sort the in-memory data using built-in sort package
	sort.Ints(inMemoryData)

	// Write the sorted data as response
	for _, value := range inMemoryData {
		fmt.Fprintf(w, "%d\n", value)
	}
}

func main() {
	http.HandleFunc("/sort", handleSortData)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
