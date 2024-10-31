package patient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatientValidation(t *testing.T) {
	t.Run("Valid patient data", func(t *testing.T) {
		p := Patient{
			ID:   "1",
			Name: "John Doe",
			Age:  30,
		}
		assert.Nil(t, p.Validate())
	})

	t.Run("Invalid patient name", func(t *testing.T) {
		p := Patient{
			ID:  "1",
			Age: 30,
		}
		assert.EqualError(t, p.Validate(), "patient name is required")
	})

	t.Run("Invalid patient age", func(t *testing.T) {
		p := Patient{
			ID:   "1",
			Name: "John Doe",
			Age:  -1,
		}
		assert.EqualError(t, p.Validate(), "patient age must be positive")
	})
}
