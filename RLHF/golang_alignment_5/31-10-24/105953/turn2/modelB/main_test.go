// patient_service_test.go

package patient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockPatientRepo struct {
	mock.Mock
}

func (m *mockPatientRepo) GetPatient(ctx context.Context, id string) (*Patient, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Patient), args.Error(1)
}

func TestPatientService_GetPatient(t *testing.T) {
	t.Run("Patient found", func(t *testing.T) {
		mockRepo := new(mockPatientRepo)
		expectedPatient := &Patient{ID: "1", Name: "John Doe", Age: 30}
		mockRepo.On("GetPatient", mock.Anything, "1").Return(expectedPatient, nil)

		service := &patientService{db: mockRepo}
		actualPatient, err := service.GetPatient(context.Background(), "1")
		assert.Nil(t, err)
		assert.Equal(t, expectedPatient, actualPatient)
		mockRepo.AssertExpectations(t)
	})
	// (Implement remaining TDD tests for other service methods)
}
