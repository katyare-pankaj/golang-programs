package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Custom structs to represent the data
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

// Custom validation function for User struct
func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid email format")
	}
	if u.Age <= 0 {
		return errors.New("age must be a positive integer")
	}
	return nil
}

// Function to parse and validate a CSV record
func (u *User) parseAndValidateCSVRecord(record []string) error {
	// We expect a record with exactly 4 fields: id, username, email, age
	if len(record) != 4 {
		return errors.Errorf("expected 4 fields in record, got: %d", len(record))
	}

	var err error
	u.ID, err = parseInt64(record[0])
	if err != nil {
		return errors.Wrapf(err, "failed to parse field 'id'")
	}

	u.Username = record[1]
	u.Email = record[2]
	u.Age, err = parseInt(record[3])
	if err != nil {
		return errors.Wrapf(err, "failed to parse field 'age'")
	}

	// Validate the user object
	if err := u.Validate(); err != nil {
		return errors.Wrap(err, "validation failed")
	}

	return nil
}

// Helper function to parse an integer (int)
func parseInt(value string) (int, error) {
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to convert '%s' to an integer", value)
	}
	return parsedValue, nil
}

// Helper function to parse an integer (int64)
func parseInt64(value string) (int64, error) {
	parsedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to convert '%s' to an int64", value)
	}
	return parsedValue, nil
}

// Function to migrate data from a CSV file to a JSON file
func migrateDataFromCSVToJSON(csvFilePath, jsonFilePath string) error {
	// Read CSV data from file
	file, err := os.Open(csvFilePath)
	if err != nil {
		return errors.Wrap(err, "failed to open CSV file")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return errors.Wrap(err, "failed to read CSV data")
	}

	// Validate CSV header and row count
	expectedHeader := []string{"id", "username", "email", "age"}
	if len(records) == 0 {
		return errors.New("CSV file is empty")
	}
	if len(records[0]) != len(expectedHeader) {
		return errors.Errorf("CSV file header doesn't match the expected format. Expected header: %s", expectedHeader)
	}
	for i, h := range records[0] {
		if h != expectedHeader[i] {
			return errors.Errorf("CSV file header field %d doesn't match the expected format. Expected: %s, Got: %s", i+1, expectedHeader[i], h)
		}
	}

	// Parse CSV data into slice of User structs and validate each row
	var users []*User
	for _, record := range records[1:] {
		user := &User{}
		if err := user.parseAndValidateCSVRecord(record); err != nil {
			log.Printf("Warning: Ignoring invalid record: %v\n", record)
			continue
		}
		users = append(users, user)
	}

	// Write JSON data to file
	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal JSON data")
	}

	err = ioutil.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		return errors.Wrap(err, "failed to write JSON file")
	}

	return nil
}

func main() {
	// Paths to the CSV and JSON files
	csvFilePath := "input.csv"
	jsonFilePath := "output.json"

	// Migrate data from CSV to JSON
	if err := migrateDataFromCSVToJSON(csvFilePath, jsonFilePath); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Data migration completed successfully!")
}
