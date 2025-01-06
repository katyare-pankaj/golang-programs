package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Displaying some HTTP status codes and their descriptions
	fmt.Println("Status Code: 200", http.StatusText(200))
	fmt.Println("Status Code: 404", http.StatusText(404))
	fmt.Println("Status Code: 500", http.StatusText(500))
	// Implementing a function that can retrieve status code description
	statusCode, err := getStatusCodeFromServer()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Server Status:", http.StatusText(statusCode))
}

func getStatusCodeFromServer() (int, error) {
	// Simulate getting a status code from a server response
	return 403, nil // Change the status code as per your requirement
}
