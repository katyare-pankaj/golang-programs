package main

import (
	"fmt"
	"net/url"
)

func main() {
	inputURL := "http://invalidurl.com"
	_, err := url.Parse(inputURL)
	if err != nil {
		fmt.Println(contextualErrorMessage(inputURL, err))
	}
}

func contextualErrorMessage(url string, err error) string {
	return fmt.Sprintf("Error parsing URL: %q\n"+
		"The provided URL is invalid. Please check the following:\n"+
		"- Ensure the URL is in the correct format (e.g., 'http://example.com').\n"+
		"- Verify the hostname is valid and reachable.\n"+
		"Underlying error: %s", url, err)
}
