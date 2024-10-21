package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

// Define a custom error type
type apiError struct {
	err    error
	status int
}

func (e *apiError) Error() string {
	return e.err.Error()
}

func (e *apiError) Status() int {
	return e.status
}

// Sample function to demonstrate error
func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, &apiError{errors.New("division by zero is not allowed"), http.StatusBadRequest}
	}
	result := dividend / divisor
	return result, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	dividendStr := r.URL.Query().Get("dividend")
	divisorStr := r.URL.Query().Get("divisor")

	// Check if dividend or divisor parameters are empty
	if dividendStr == "" || divisorStr == "" {
		handleError(w, r, &apiError{errors.New("'dividend' or 'divisor' parameter is missing"), http.StatusBadRequest})
		return
	}

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
		return 0, &apiError{errors.Wrap(err, "error parsing float"), http.StatusBadRequest}
	}
	return value, nil
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	apiErr, ok := err.(*apiError)
	if !ok {
		apiErr = &apiError{err, http.StatusInternalServerError}
	}

	http.Error(w, apiErr.Error(), apiErr.Status())
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
