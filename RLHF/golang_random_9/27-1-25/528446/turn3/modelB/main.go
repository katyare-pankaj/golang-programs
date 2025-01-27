package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/mail"
	"os"
	"strings"
)

// Custom error type to track error origin
type EmailError struct {
	Function string
	Err      error
}

func (e *EmailError) Error() string {
	return fmt.Sprintf("error in %s: %v", e.Function, e.Err)
}

// Function to validate email format
func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

// Function to read input email addresses from user and parse them
func readEmails() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter email addresses separated by commas: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, &EmailError{Function: "readEmails", Err: err}
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return nil, &EmailError{Function: "readEmails", Err: errors.New("input is empty")}
	}
	emails := strings.Split(input, ",")
	var validEmails []string

	for _, email := range emails {
		email = strings.TrimSpace(email)
		err := validateEmail(email)
		if err != nil {
			return nil, &EmailError{Function: "readEmails", Err: fmt.Errorf("invalid email format: %s", email)}
		}
		validEmails = append(validEmails, email)
	}

	return validEmails, nil
}

func main() {
	emails, err := readEmails()
	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}

	fmt.Println("Valid email addresses:")
	for _, email := range emails {
		fmt.Println(email)
	}
}
