package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// database cache ,don't use this in real world.just to see
var databaseData = make(map[int]string)

func databaseGetHandler(w http.ResponseWriter, r *http.Request) {
	//validate the JWT token in the request header here
	//Assume the user is authenticated using a valid JWT token for simplicity
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	data := databaseData[id]
	if data == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, data)
}

func main() {
	r := mux.NewRouter()
	//Access Control Level 1:
	//Public Endpoints
	r.HandleFunc("/ping", pingHandler).Methods("GET")

	//Access Control Level 2:
	//Authantiation mandatory
	r.HandleFunc("/api/data/{id}", databaseGetHandler).Methods("GET")

	// Initialize SSL certificates in real world.
	http.ListenAndServe(":8080", r)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pong!")
}
