package main

import (
	"fmt"
	"regexp"
)

// AnonymizeSensitiveData takes a slice of strings containing sensitive user data and anonymizes emails, phones, credit cards, and addresses
func AnonymizeSensitiveData(data []string) []string {
	emailPattern := regexp.MustCompile(`[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,}`)
	phonePattern := regexp.MustCompile(`\+?[0-9]{1,4}[ -]?$?[0-9]{1,3}?$?[ -]?[0-9]{3,4}[ -]?[0-9]{3,4}`)
	creditCardPattern := regexp.MustCompile(`\b(?:\d[ -]*?){13,16}\b`)
	addressPattern := regexp.MustCompile(`\d{1,5} [\w\s]{2,30}(?:Street|St|Avenue|Ave|Boulevard|Blvd|Rd|Road|Lane|Ln|Way)(?:,?\s+\w+\s*\d{5})?`)

	anonymizedData := make([]string, len(data))
	copy(anonymizedData, data)

	for i, value := range anonymizedData {
		anonymizedValue := emailPattern.ReplaceAllString(value, "user@example.com")
		anonymizedValue = phonePattern.ReplaceAllString(anonymizedValue, "xxx-xxx-xxxx")
		anonymizedValue = creditCardPattern.ReplaceAllString(anonymizedValue, "xxxx-xxxx-xxxx-xxxx")
		anonymizedValue = addressPattern.ReplaceAllString(anonymizedValue, "#### Anonymized Address ####")

		anonymizedData[i] = anonymizedValue
	}

	return anonymizedData
}

func main() {
	data := []string{
		"Contact: john.doe@example.com | Phone: 123-456-7890 | Card: 4111-1111-1111-1111 | Address: 123 Elm Street, Springfield 12345",
		"Email: jane.doe@example.org | Phone: +1 (555) 123-4567 | Card: 5500 0000 0000 0004 | Address: 456 Oak Ave, Metropolis, 67890",
	}

	anonymizedData := AnonymizeSensitiveData(data)

	fmt.Println("Original Data:")
	for _, v := range data {
		fmt.Println(v)
	}

	fmt.Println("\nAnonymized Data:")
	for _, v := range anonymizedData {
		fmt.Println(v)
	}
}
