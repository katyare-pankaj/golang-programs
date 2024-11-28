package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ErrUnauthorized struct {
	message string
}

func (e *ErrUnauthorized) Error() string {
	return e.message
}

func NewErrUnauthorized(message string) error {
	return &ErrUnauthorized{message: message}
}

type ErrForbidden struct {
	message string
}

func (e *ErrForbidden) Error() string {
	return e.message
}

func NewErrForbidden(message string) error {
	return &ErrForbidden{message: message}
}

func MyFunction(role string) error {
	if role != "admin" {
		return NewErrUnauthorized("You are not authorized to perform this action")
	}

	// Some other logic here

	return nil
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("X-User-Role")

	err := MyFunction(role)
	if err != nil {
		switch err.(type) {
		case *ErrUnauthorized:
			http.Error(w, err.Error(), http.StatusUnauthorized)
		case *ErrForbidden:
			http.Error(w, err.Error(), http.StatusForbidden)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Handle success
	w.Write([]byte("success"))
}
func main() {
	route := mux.NewRouter()

	route.HandleFunc("/api", MyHandler)

	fmt.Print("server listening code 8080")

	http.ListenAndServe(":8080", route)
}
