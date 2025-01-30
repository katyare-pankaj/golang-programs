package main

import (
	"fmt"
	"regexp"
)

// AnonymizeSensitiveData anonymizes sensitive data in a slice of strings
func AnonymizeSensitiveData(data []string) []string {
	// Regular expression patterns for matching sensitive data
	emailPattern := regexp.MustCompile(`[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,}`)
	phonePattern := regexp.MustCompile(`\+?[0-9]{1,4}[ -]?$?[0-9]{1,3}?$?[ -]?[0-9]{3,4}[ -]?[0-9]{3,4}`)
	creditCardPattern := regexp.MustCompile(`(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})`)
	addressPattern := regexp.MustCompile(`\b\d{1,5} [a-zA-Z0-9 ]+\b`)

	// Create an anonymized copy of the data slice
	anonymizedData := make([]string, len(data))
	copy(anonymizedData, data)

	// Anonymize each part of the data
	for i, value := range anonymizedData {
		// Replace emails with a placeholder
		anonymizedValue := emailPattern.ReplaceAllString(value, "user@example.com")

		// Replace phone numbers with a placeholder
		anonymizedValue = phonePattern.ReplaceAllString(anonymizedValue, "xxx-xxx-xxxx")

		// Replace credit card numbers with a placeholder
		anonymizedValue = creditCardPattern.ReplaceAllString(anonymizedValue, "xxxxxxxxxxxxxxxx")

		// Replace addresses with a placeholder
		anonymizedValue = addressPattern.ReplaceAllString(anonymizedValue, "xxx xxxx xxxx")

		// Update the anonymized value in the slice
		anonymizedData[i] = anonymizedValue
	}

	return anonymizedData
}

func main() {
	// Example dataset with sensitive data
	data := []string{
		"John Doe john.doe@example.com 123-456-7890 4111111111111111 123 Main St, Anytown, USA",
		"Jane Doe jane.doe@example.org +1 (555) 123-4567 5555-5555-5555-5555 456 Elm St, Cityville, CA",
		"Alice alice@example.net 555-456-1234 378282246310005 789 Oak St, Suburbia, TX",
	}

	// Anonymize the data
	anonymizedData := AnonymizeSensitiveData(data)

	// Print results
	fmt.Println("Original Data:")
	for _, v := range data {
		fmt.Println(v)
	}

	fmt.Println("\nAnonymized Data:")
	for _, v := range anonymizedData {
		fmt.Println(v)
	}
}
