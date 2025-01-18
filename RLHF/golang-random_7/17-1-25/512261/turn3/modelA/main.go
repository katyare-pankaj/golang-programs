package main

import (
	"fmt"
	"io"
	"net/http"
)

func thirdPartyFunction(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Simulate an error scenario that could lead to a panic
	if url == "https://example.com/panic" {
		panic("Third-party function intentionally panics!")
	}

	return res, nil
}

func main() {
	urls := []string{
		"https://example.com",
		"https://example.com/panic",
	}

	for _, url := range urls {
		fmt.Println("Calling third-party function with URL:", url)

		if res, err := thirdPartyFunction(url); err != nil {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
				continue
			}
			fmt.Println("Error:", err)
		} else {
			defer res.Body.Close()
			io.Copy(io.Discard, res.Body)
			fmt.Println("Success.")
		}
	}
}
