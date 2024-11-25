package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/25-11-24/360575/turn1/modelA/envutil"
	"log"
	"net/http"
)

func main() {

	port := envutil.GetEnv("PORT", "8080")
	fmt.Printf("Starting server on port: %s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\n")
	})

	log.Fatalf("Failed to start server: %v", http.ListenAndServe(":"+port, nil))
}
