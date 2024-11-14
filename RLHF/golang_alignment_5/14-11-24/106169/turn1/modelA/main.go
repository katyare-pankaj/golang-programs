package main

import (
	"fmt"
	"regexp"
)

func scrubEmails(emails []string) []string {
	scrubbed := make([]string, 0)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

	for _, email := range emails {
		if emailRegex.MatchString(email) {
			scrubbed = append(scrubbed, email)
		}
	}
	return scrubbed
}

func main() {
	emails := []string{
		"example@example.com",
		"invalidemail",
		"another@example.com",
		"name with space@example.com", // Invalid according to the simple regex
	}
	scrubbed := scrubEmails(emails)
	fmt.Println("Scrubbed emails:", scrubbed)
}
