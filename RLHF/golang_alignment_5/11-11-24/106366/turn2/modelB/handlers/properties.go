// handlers/properties.go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/property-management/api/models"
	"github.com/property-management/api/services"
)

// CreatePropertyHandler handles the creation of a new property
func CreatePropertyHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a Property struct
	var property models.Property
	if err := json.NewDecoder(r.Body).Decode(&property); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the property using the PropertyService
	if err := services.PropertyService.Create(&property); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response status to 201 Created
	w.WriteHeader(http.StatusCreated)

	// Encode the created property in the response body
	if err := json.NewEncoder(w).Encode(&property); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
