package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// store to store idempotency-key data
var storage = make(map[string]string)
var mu sync.Mutex

// idempotentHandler is a handler function which returns a 200 OK response if the request is idempotent or 409 Conflict if the request is not idempotent.
func idempotentHandler(w http.ResponseWriter, r *http.Request) {
	//Get the 'Idempotency-Key' header from the request
	idempotencyKey := r.Header.Get("Idempotency-Key")

	//Acquire lock to access the shared storage
	mu.Lock()
	defer mu.Unlock()

	//Check if the idempotency-key is already present in the storage
	if _, ok := storage[idempotencyKey]; ok {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Duplicate request detected\n")
		return
	}
	//store the idempotency-key in the storage
	storage[idempotencyKey] = "Processed"
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Request processed successfully\n")
}

func main() {
	//Register the idempotentHandler function for the "/idempotent" route
	http.HandleFunc("/idempotent", idempotentHandler)

	//Start the HTTP server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
