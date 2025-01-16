package main

import (
	"fmt"
	"html"
	"net/http"
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
	escapedInput := html.EscapeString(userInput)
	output := fmt.Sprintf("Welcome, %s!", escapedInput)
	w.Write([]byte(output))
}
