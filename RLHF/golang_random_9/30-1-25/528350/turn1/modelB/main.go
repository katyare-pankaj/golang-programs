package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Library functions for anonymizing sensitive data

// AnonymizeStrings takes a slice of strings and returns an anonymized slice.
func AnonymizeStrings(data []string) []string {
	anonymizedData := make([]string, len(data))
	rand.Seed(time.Now().UnixNano())

	for i, value := range data {
		// You can customize the anonymization process here.
		// For simplicity, this example uses random strings.
		anonymizedData[i] = generateRandomString(10)
	}
	return anonymizedData
}

// AnonymizeEmails takes a slice of email strings and returns an anonymized slice.
func AnonymizeEmails(emails []string) []string {
	anonymizedEmails := make([]string, len(emails))
	for i, email := range emails {
		// Split the email into local part and domain
		localPart, domain := splitEmail(email)
		// Anonymize the local part
		anonymizedLocalPart := generateRandomString(8)
		// Combine the anonymized local part and domain
		anonymizedEmail := anonymizedLocalPart + "@" + domain
		anonymizedEmails[i] = anonymizedEmail
	}
	return anonymizedEmails
}

// Private helper functions

func generateRandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func splitEmail(email string) (string, string) {
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		return "", "" // Invalid email, return empty strings
	}
	return email[:atIndex], email[atIndex+1:]
}

// Main function to test the library
func main() {
	// Sample sensitive data
	sensitiveData := []string{"Alice", "Bob", "Charlie", "David", "Emma"}
	emails := []string{"alice@example.com", "bob@example.org", "charlie@gmail.com"}

	// Anonymize the data
	anonymizedData := AnonymizeStrings(sensitiveData)
	anonymizedEmails := AnonymizeEmails(emails)

	// Display the original and anonymized data
	fmt.Println("Original Data:", sensitiveData)
	fmt.Println("Anonymized Data:", anonymizedData)
	fmt.Println("Original Emails:", emails)
	fmt.Println("Anonymized Emails:", anonymizedEmails)
}
