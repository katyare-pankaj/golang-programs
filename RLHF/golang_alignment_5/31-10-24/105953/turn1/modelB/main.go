// patient.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Patient struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	BloodType string `json:"bloodType"`
}

var patients []Patient

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/patients", getPatients).Methods("GET")
	r.HandleFunc("/patients", addPatient).Methods("POST")
	r.HandleFunc("/patients/{id}", getPatient).Methods("GET")
	r.HandleFunc("/patients/{id}", updatePatient).Methods("PUT")
	r.HandleFunc("/patients/{id}", deletePatient).Methods("DELETE")

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// getPatients handles GET requests to /patients
func getPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}

// addPatient handles POST requests to /patients
func addPatient(w http.ResponseWriter, r *http.Request) {
	var patient Patient
	_ = json.NewDecoder(r.Body).Decode(&patient)
	patients = append(patients, patient)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

// getPatient handles GET requests to /patients/{id}
func getPatient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for _, patient := range patients {
		if patient.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(patient)
			return
		}
	}
	http.Error(w, "Patient not found", http.StatusNotFound)
}

// updatePatient handles PUT requests to /patients/{id}
func updatePatient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var patient Patient
	_ = json.NewDecoder(r.Body).Decode(&patient)
	for i, p := range patients {
		if p.ID == id {
			patients[i] = patient
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(patient)
			return
		}
	}
	http.Error(w, "Patient not found", http.StatusNotFound)
}

// deletePatient handles DELETE requests to /patients/{id}
func deletePatient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for i, patient := range patients {
		if patient.ID == id {
			patients = append(patients[:i], patients[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Patient not found", http.StatusNotFound)
}
