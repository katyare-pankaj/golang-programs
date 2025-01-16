package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	userInput := r.FormValue("userInput")
	escapedInput := escapeString(userInput)
	output := fmt.Sprintf("Welcome, %s!", escapedInput)
	w.Write([]byte(output))
}

func escapeString(input string) string {
	// Replace single quotes with double quotes
	input = strings.ReplaceAll(input, "'", `"`)
	// Replace double quotes with escaped double quotes
	input = strings.ReplaceAll(input, `"`, `\"`)
	// Replace backslashes with escaped backslashes
	input = strings.ReplaceAll(input, `\`, `\\`)
	// Replace ampersands with escaped ampersands
	input = strings.ReplaceAll(input, `&`, `&amp;`)
	// Replace less than signs with escaped less than signs
	input = strings.ReplaceAll(input, `<`, `&lt;`)
	// Replace greater than signs with escaped greater than signs
	input = strings.ReplaceAll(input, `>`, `&gt;`)
	// Replace copyright symbol with escaped copyright symbol
	input = strings.ReplaceAll(input, `©`, `&copy;`)
	// Replace registered trademark symbol with escaped registered trademark symbol
	input = strings.ReplaceAll(input, `®`, `&reg;`)
	return input
}
