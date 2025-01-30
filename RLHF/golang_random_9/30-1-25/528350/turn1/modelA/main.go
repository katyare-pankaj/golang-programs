package main

import (
	"fmt"
	"regexp"
)

// AnonymizeData takes a slice of strings representing sensitive user data and returns a new slice with data anonymized
func AnonymizeData(data []string) []string {
	emailPattern := regexp.MustCompile(`[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,}`)
	phonePattern := regexp.MustCompile(`\+?[0-9]{1,4}[ -]?$?[0-9]{1,3}?$?[ -]?[0-9]{3,4}[ -]?[0-9]{3,4}`)

	anonymizedData := make([]string, len(data))
	copy(anonymizedData, data)

	for i, value := range anonymizedData {
		anonymizedValue := emailPattern.ReplaceAllString(value, "anonymous@example.com")
		anonymizedValue = phonePattern.ReplaceAllString(anonymizedValue, "xxx-xxx-xxxx")
		anonymizedData[i] = anonymizedValue
	}

	return anonymizedData
}

func main() {
	// Example dataset containing email addresses and phone numbers
	data := []string{
		"John Doe john.doe@example.com 123-456-7890",
		"Jane Doe jane.doe@example.org +1 (555) 123-4567",
		"Alice alice@example.net 555-456-1234",
	}

	anonymizedData := AnonymizeData(data)

	fmt.Println("Original Data:")
	for _, v := range data {
		fmt.Println(v)
	}

	fmt.Println("\nAnonymized Data:")
	for _, v := range anonymizedData {
		fmt.Println(v)
	}
}
