package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	response := fmt.Sprintf("Hello, my name is %s and I am %s years old.", name, age)
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/greet", handler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
