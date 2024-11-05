// patient_service.go

package patient

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Patient represents a patient entity
type Patient struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	HealthStatus string `json:"health_status"`
}

// PatientService interface defines the service methods
type PatientService interface {
	GetPatient(ctx context.Context, id string) (*Patient, error)
	CreatePatient(ctx context.Context, patient *Patient) (*Patient, error)
	UpdatePatient(ctx context.Context, patient *Patient) (*Patient, error)
	DeletePatient(ctx context.Context, id string) error
}

// patientService implements the PatientService interface
type patientService struct {
	db *gorm.DB
}

// GetPatient implements the PatientService interface
func (s *patientService) GetPatient(ctx context.Context, id string) (*Patient, error) {
	var patient Patient
	if err := s.db.Where("id = ?", id).First(&patient).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("patient with id %s not found", id)
		}
		return nil, err
	}
	return &patient, nil
}

// (Implement remaining service methods and TDD tests)
