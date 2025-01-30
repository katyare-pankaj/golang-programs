package main

import (
	"fmt"
	"regexp"
)

// AnonymizeEmailsAndPhones anonymizes emails and phone numbers in a slice of strings
func AnonymizeEmailsAndPhones(data []string) []string {
	// Regular expression patterns for matching emails and phone numbers
	emailPattern := regexp.MustCompile(`[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,}`)
	phonePattern := regexp.MustCompile(`\+?[0-9]{1,4}[ -]?$?[0-9]{1,3}?$?[ -]?[0-9]{3,4}[ -]?[0-9]{3,4}`)

	// Create an anonymized copy of the data slice
	anonymizedData := make([]string, len(data))
	copy(anonymizedData, data)

	// Anonymize each part of the data
	for i, value := range anonymizedData {
		// Replace emails with a placeholder
		anonymizedValue := emailPattern.ReplaceAllString(value, "user@example.com")

		// Replace phone numbers with a placeholder
		anonymizedValue = phonePattern.ReplaceAllString(anonymizedValue, "xxx-xxx-xxxx")

		// Update the anonymized value in the slice
		anonymizedData[i] = anonymizedValue
	}

	return anonymizedData
}

// Example of using the AnonymizeEmailsAndPhones function
func main() {
	// Example dataset with emails and phone numbers
	data := []string{
		"Contact John at john.doe@example.com or call 123-456-7890",
		"Jane's email is jane.doe@example.org and phone is +1 (555) 123-4567",
		"Alice: alice@example.net, Phone: 555-456-1234",
	}

	// Anonymize the data
	anonymizedData := AnonymizeEmailsAndPhones(data)

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
