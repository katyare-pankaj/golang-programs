package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define our custom error type
type InvalidRequestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *InvalidRequestError) Error() string {
	return fmt.Sprintf("InvalidRequestError: %s", e.Message)
}

// We'll use this function to handle invalid requests
func handleInvalidRequest(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	// Convert the custom error to JSON and write it to the response
	json.NewEncoder(w).Encode(&InvalidRequestError{Message: err.Error(), Status: http.StatusBadRequest})
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	dividendStr := r.URL.Query().Get("dividend")
	divisorStr := r.URL.Query().Get("divisor")

	// Validate the request parameters
	if dividendStr == "" || divisorStr == "" {
		err := &InvalidRequestError{Message: "Missing required parameters: dividend and divisor", Status: http.StatusBadRequest}
		handleInvalidRequest(w, r, err)
		return
	}

	dividend, err := strconv.ParseFloat(dividendStr, 64)
	if err != nil {
		handleInvalidRequest(w, r, err)
		return
	}

	divisor, err := strconv.ParseFloat(divisorStr, 64)
	if err != nil {
		handleInvalidRequest(w, r, err)
		return
	}

	// Rest of the handler logic
	result := dividend / divisor
	fmt.Fprintf(w, "Result: %.2f\n", result)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/divide", divideHandler)

	log.Println("API server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
