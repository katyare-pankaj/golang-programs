package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func unsafeThirdPartyFunction(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func safeThirdPartyFunction(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Recover from panic and ensure to return an error
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	// Simulate a panic
	panic("Oh no! A panic occurred!")
	return string(data), nil
}

func main() {
	urls := []string{
		"https://example.com/example.txt",
		"https://example.com/non-existent-page",
	}

	for _, url := range urls {
		fmt.Println("Calling unsafeThirdPartyFunction with URL:", url)
		result, err := unsafeThirdPartyFunction(url)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}

		fmt.Println("Calling safeThirdPartyFunction with URL:", url)
		result, err = safeThirdPartyFunction(url)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
		fmt.Println()
	}
}
