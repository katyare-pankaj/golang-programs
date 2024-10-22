package main

import (
	"regexp"
	"testing"
)

// employee represents an employee record.
type employee struct {
	firstName string
	lastName  string
	email     string
}

// isValidEmployeeData checks the consistency of the employee data.
func isValidEmployeeData(e employee) bool {
	if e.firstName == "" || e.lastName == "" {
		return false
	}

	// Regular expression for validating an email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(e.email) {
		return false
	}

	return true
}

func TestEmployeeDataConsistency(t *testing.T) {
	testCases := []struct {
		name     string
		employee employee
		expected bool
	}{
		{
			name:     "Valid Employee Data",
			employee: employee{firstName: "John", lastName: "Doe", email: "johndoe@example.com"},
			expected: true,
		},
		{
			name:     "Invalid Email Format",
			employee: employee{firstName: "Jane", lastName: "Smith", email: "janesmith"},
			expected: false,
		},
		{
			name:     "Empty First Name",
			employee: employee{firstName: "", lastName: "Doe", email: "johndoe@example.com"},
			expected: false,
		},
		{
			name:     "Empty Last Name",
			employee: employee{firstName: "John", lastName: "", email: "johndoe@example.com"},
			expected: false,
		},
		{
			name:     "Empty Email",
			employee: employee{firstName: "John", lastName: "Doe", email: ""},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isValidEmployeeData(tc.employee)
			if actual != tc.expected {
				t.Errorf("Expected %t, got %t", tc.expected, actual)
			}
		})
	}
}
