package main

import (
	"fmt"
	"os"
	"strings"
)

// OldContact represents a contact in the old format
type OldContact struct {
	Name  string
	Email string
}

// NewContact represents a contact in the new format
type NewContact struct {
	FirstName string
	LastName  string
	Email     string
}

// ValidateOldContact ensures the old contact data is valid
func ValidateOldContact(c OldContact) error {
	if len(c.Name) == 0 {
		return fmt.Errorf("name is required")
	}
	if !strings.Contains(c.Email, "@") {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

// ValidateNewContact ensures the new contact data is valid
func ValidateNewContact(c NewContact) error {
	if len(c.FirstName) == 0 || len(c.LastName) == 0 {
		return fmt.Errorf("first name and last name are required")
	}
	if !strings.Contains(c.Email, "@") {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

// MigrateContacts migrates a list of contacts from old to new format
func MigrateContacts(oldContacts []OldContact) ([]NewContact, error) {
	var newContacts []NewContact
	for _, oldContact := range oldContacts {
		err := ValidateOldContact(oldContact)
		if err != nil {
			return nil, fmt.Errorf("validation failed for old contact: %v", err)
		}

		// Simple transformation: split name into first and last
		nameParts := strings.Split(oldContact.Name, " ")
		if len(nameParts) != 2 {
			return nil, fmt.Errorf("cannot split name: %s", oldContact.Name)
		}
		newContact := NewContact{
			FirstName: nameParts[0],
			LastName:  nameParts[1],
			Email:     oldContact.Email,
		}

		err = ValidateNewContact(newContact)
		if err != nil {
			return nil, fmt.Errorf("validation failed for new contact: %v", err)
		}

		newContacts = append(newContacts, newContact)
	}

	return newContacts, nil
}

// WriteContacts writes the new contacts to a file
func WriteContacts(contacts []NewContact, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	for _, contact := range contacts {
		_, err = file.WriteString(fmt.Sprintf("%s, %s, %s\n", contact.FirstName, contact.LastName, contact.Email))
		if err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}
	}

	return nil
}

func main() {
	// Example old contact data
	oldContacts := []OldContact{
		{Name: "John Doe", Email: "john@example.com"},
		{Name: "Jane Smith", Email: "jane@example.com"},
	}

	// Migrate contacts
	newContacts, err := MigrateContacts(oldContacts)
	if err != nil {
		fmt.Println("Data migration failed:", err)
		return
	}

	// Write new contacts to file
	outputPath := "new_contacts.csv"
	err = WriteContacts(newContacts, outputPath)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data migration successful. Results written to", outputPath)
}
