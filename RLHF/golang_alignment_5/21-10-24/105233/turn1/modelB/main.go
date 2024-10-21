package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// CustomError represents a custom error type with a message
type CustomError struct {
	message string
}

func (e CustomError) Error() string {
	return e.message
}

// readFile reads the content of a file and returns it as a string
func readFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", &CustomError{message: "Error reading file: " + err.Error()}
	}
	return string(data), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	if filename == "" {
		http.Error(w, "Invalid request: 'file' parameter is missing", http.StatusBadRequest)
		return
	}

	content, err := readFile(filename)
	if err != nil {
		// Handle custom errors
		if customErr, ok := err.(*CustomError); ok {
			http.Error(w, customErr.message, http.StatusInternalServerError)
			return
		}

		// Handle other generic errors
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, content)
}

func main() {
	http.HandleFunc("/readfile", handler)
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
