package main

import (
	"fmt"
	"io"
	"net/http"
)

func unsafeThirdPartyFunction(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res, nil
}

func safeThirdPartyFunction(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			defer res.Body.Close()
			if err, ok := r.(error); ok {
				fmt.Println("Recovered error: ", err)
			} else {
				panic(r)
			}
		}
	}()

	// Simulate error
	if url == "https://example.com/bad-url" {
		return nil, fmt.Errorf("bad url")
	}

	return res, nil
}

func main() {
	urls := []string{
		"https://example.com",
		"https://example.com/bad-url",
	}

	for _, url := range urls {
		fmt.Println("Calling unsafeThirdPartyFunction with URL:", url)
		res, err := unsafeThirdPartyFunction(url)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			defer res.Body.Close()
			io.Copy(io.Discard, res.Body)
			fmt.Println("Success.")
		}

		fmt.Println("Calling safeThirdPartyFunction with URL:", url)
		res, err = safeThirdPartyFunction(url)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			defer res.Body.Close()
			io.Copy(io.Discard, res.Body)
			fmt.Println("Success.")
		}
		fmt.Println()
	}
}
