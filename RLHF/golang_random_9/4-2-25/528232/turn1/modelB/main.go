package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func main() {
	apiURL := "https://example.com/api/data" // Replace this with your third-party API endpoint

	data, err := getDataFromAPI(apiURL)
	if err != nil {
		log.Fatalf("Error getting data from API: %v", err)
	}

	fmt.Println(data)
}

func getDataFromAPI(apiURL string) (string, error) {
	// Create a new RESTy client
	client := resty.New()

	// Use defer to close the connection after the request is made
	defer client.Close()

	// Make a GET request to the API endpoint
	resp, err := client.R().Get(apiURL)
	if err != nil {
		return "", err
	}

	// Check the status code
	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("API request failed with status: %s", resp.Status())
	}

	// Get the response body as a string
	data := string(resp.Body())

	return data, nil
}
