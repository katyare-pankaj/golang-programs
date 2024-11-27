package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/example", exampleHandler).Methods("GET")

	fmt.Println("On-premises API gateway running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, this is an example API response!")
}
