package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Define a custom error type
type apiError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *apiError) Error() string {
	return fmt.Sprintf("api error: %s (status %d)", e.Message, e.Status)
}

// Sample function that might encounter errors
func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, &apiError{Message: "division by zero is not allowed", Status: http.StatusBadRequest}
	}
	result := dividend / divisor
	return result, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	dividendStr := r.URL.Query().Get("dividend")
	divisorStr := r.URL.Query().Get("divisor")

	dividend, err := parseFloat(dividendStr)
	if err != nil {
		handleError(w, r, err)
		return
	}

	divisor, err := parseFloat(divisorStr)
	if err != nil {
		handleError(w, r, err)
		return
	}

	result, err := divide(dividend, divisor)
	if err != nil {
		handleError(w, r, err)
		return
	}

	fmt.Fprintf(w, "Result: %.2f\n", result)
}

func parseFloat(str string) (float64, error) {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, &apiError{Message: "error parsing float", Status: http.StatusBadRequest}
	}
	return value, nil
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	apiErr, ok := err.(*apiError)
	if !ok {
		apiErr = &apiError{Message: "internal server error", Status: http.StatusInternalServerError}
	}

	http.Error(w, apiErr.Error(), apiErr.Status)
}

func main() {
	r := chi.NewRouter()
	r.Get("/divide", handler)

	fmt.Println("API server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
