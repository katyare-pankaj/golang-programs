// main.go
package main

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/08-11-24/106485/turn2/modelB/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Register the routes

	router.HandleFunc("/properties", handlers.CreatePropertyHandler).Methods("POST")
	router.HandleFunc("/properties/{id}", handlers.GetPropertyHandler).Methods("GET")
	router.HandleFunc("/properties/{id}", handlers.UpdatePropertyHandler).Methods("PUT")
	router.HandleFunc("/properties/{id}", handlers.DeletePropertyHandler).Methods("DELETE")

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
