package main

import (
	"go-programs/RLHF/golang_alignment_5/31-10-24/105953/turn2/modelA/patient"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	patientHandler := patient.NewPatientHandler()
	patientHandler.RegisterRoutes(router)

	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
