// main.go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")

	//Start monitoring performance
	startTime := time.Now()

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Failed to start the server: %v", err)
	}

	// End performance monitoring
	elapsedTime := time.Since(startTime)
	fmt.Printf("Application performance time: %s\n", elapsedTime)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
