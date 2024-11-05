// patienthandler.go

package patient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Patient represents a patient entity
type Patient struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	HealthStatus string `json:"health_status"`
}

// PatientHandler handles patient-related requests
type PatientHandler struct{}

// NewPatientHandler creates a new PatientHandler
func NewPatientHandler() *PatientHandler {
	return &PatientHandler{}
}

// RegisterRoutes registers the patient-related routes with the router
func (h *PatientHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/patients", h.handleGetPatients).Methods("GET")
	r.HandleFunc("/patients/{id}", h.handleGetPatient).Methods("GET")
}

// handleGetPatients handles the GET /patients request
func (h *PatientHandler) handleGetPatients(w http.ResponseWriter, r *http.Request) {
	// Implementation for retrieving multiple patients goes here
	// For demonstration, we'll just return a stubbed response
	patients := []Patient{{ID: "1", Name: "Patient A"}, {ID: "2", Name: "Patient B"}}
	json.NewEncoder(w).Encode(patients)
}

// handleGetPatient handles the GET /patients/{id} request
func (h *PatientHandler) handleGetPatient(w http.ResponseWriter, r *http.Request) {
	// Implementation for retrieving a single patient goes here
	// For demonstration, we'll just return a stubbed response
	vars := mux.Vars(r)
	patientID := vars["id"]
	patient := Patient{ID: patientID, Name: fmt.Sprintf("Patient %s", patientID)}
	json.NewEncoder(w).Encode(patient)
}

// patient.go (remains the same)
