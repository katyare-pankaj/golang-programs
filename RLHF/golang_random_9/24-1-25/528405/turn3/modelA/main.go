package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// CustomError is a type that includes additional context about an error
type CustomError struct {
	Function string
	Err      error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error in %s: %v", e.Function, e.Err)
}

// readInput prompts the user for input and returns the input along with any error
func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a list of email addresses separated by commas: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", &CustomError{Function: "readInput", Err: err}
	}
	text = strings.TrimSpace(text)
	return text, nil
}

// validateEmails splits the input into emails and validates each email
func validateEmails(input string) ([]string, error) {
	if input == "" {
		return nil, &CustomError{Function: "validateEmails", Err: errors.New("input is empty")}
	}

	emails := strings.Split(input, ",")
	validEmails := []string{}
	invalidEmails := []string{}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)

	for _, email := range emails {
		email = strings.TrimSpace(email)
		if re.MatchString(email) {
			validEmails = append(validEmails, email)
		} else {
			invalidEmails = append(invalidEmails, email)
		}
	}

	if len(invalidEmails) > 0 {
		return validEmails, &CustomError{Function: "validateEmails", Err: fmt.Errorf("invalid emails: %v", invalidEmails)}
	}

	return validEmails, nil
}

func main() {
	input, err := readInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	emails, err := validateEmails(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Valid emails:", emails)
}
