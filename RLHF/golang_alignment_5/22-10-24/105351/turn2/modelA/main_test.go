package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestEmployeeDataEntry(t *testing.T) {
	type employeeDataEntry struct {
		fieldName       string
		input           string
		expectedIsValid bool
		expectedMessage string
	}

	testCases := []employeeDataEntry{
		{
			fieldName:       "First Name",
			input:           "Alice",
			expectedIsValid: true,
			expectedMessage: "",
		},
		{
			fieldName:       "First Name",
			input:           "",
			expectedIsValid: false,
			expectedMessage: "First name is required.",
		},
		{
			fieldName:       "Last Name",
			input:           "Brown",
			expectedIsValid: true,
			expectedMessage: "",
		},
		{
			fieldName:       "Email",
			input:           "alice@example",
			expectedIsValid: false,
			expectedMessage: "Invalid email format. Please enter a valid email address.",
		},
		// Add more test cases for different fields
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Validating %s", tc.fieldName), func(t *testing.T) {
			isValid, message := validateEmployeeDataEntry(tc.fieldName, tc.input)
			if isValid != tc.expectedIsValid {
				t.Errorf("Expected validation to be %t for '%s', got %t", tc.expectedIsValid, tc.input, isValid)
			}
			if message != tc.expectedMessage {
				t.Errorf("Expected message '%s', got '%s'", tc.expectedMessage, message)
			}
		})
	}
}

func validateEmployeeDataEntry(fieldName, input string) (bool, string) {
	var isValid bool
	var message string

	switch fieldName {
	case "First Name", "Last Name":
		isValid = len(strings.TrimSpace(input)) > 0
		if !isValid {
			message = fmt.Sprintf("%s is required.", fieldName)
		}
	case "Email":
		isValid = isValidEmail(input)
		if !isValid {
			message = "Invalid email format. Please enter a valid email address."
		}
	// Add more cases for other fields with specific validation rules
	default:
		isValid = false
		message = fmt.Sprintf("Invalid field name: %s", fieldName)
	}

	return isValid, message
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
	return strings.TrimSpace(email) != ""
}
