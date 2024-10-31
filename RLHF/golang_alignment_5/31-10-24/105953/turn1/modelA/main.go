// patient.go

package patient

import (
	"errors"
)

// Patient represents a patient entity
type Patient struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	HealthStatus string `json:"health_status"`
}

// Validate ensures the patient data is valid
func (p *Patient) Validate() error {
	if p.Name == "" {
		return errors.New("patient name is required")
	}
	if p.Age <= 0 {
		return errors.New("patient age must be positive")
	}
	return nil
}
