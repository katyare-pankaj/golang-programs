package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Define a structure that represents a row in the CSV file
type Record struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
	Department string `json:"department"`
}

// ValidateRecord ensures the data in the record is valid
func ValidateRecord(record Record) error {
	if record.ID == 0 {
		return fmt.Errorf("ID is required")
	}
	if len(record.Name) == 0 {
		return fmt.Errorf("Name is required")
	}
	if !strings.Contains(record.Email, "@") {
		return fmt.Errorf("Invalid email format")
	}
	if record.Age < 0 {
		return fmt.Errorf("Age must be non-negative")
	}
	if len(record.Department) == 0 {
		return fmt.Errorf("Department is required")
	}
	return nil
}

// MigrateCSVToJSON migrates data from a CSV file to a JSON file
func MigrateCSVToJSON(csvPath, jsonPath string) error {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf("error opening CSV file: %v", err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	// Read the header from the CSV file
	header, err := csvReader.Read()
	if err != nil {
		return fmt.Errorf("error reading CSV header: %v", err)
	}

	// Map CSV header to record fields
	recordMap := map[string]string{}
	for index, field := range header {
		recordMap[field] = fmt.Sprintf("%d", index)
	}

	// Open the JSON file for writing
	jsonFile, err := os.Create(jsonPath)
	if err != nil {
		return fmt.Errorf("error creating JSON file: %v", err)
	}
	defer jsonFile.Close()

	jsonEncoder := json.NewEncoder(jsonFile)
	jsonEncoder.SetIndent("", "  ")

	var records []Record
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("error reading CSV row: %v", err)
		}

		// Create a new record and populate it from the CSV row
		record := Record{}
		for _, field := range header {
			valueIndex := recordMap[field]
			if valueIndex == "" {
				continue // Skip fields not in the header
			}
			switch field {
			case "ID", "Age":
				var intValue int
				if err := fmt.Sscanf(row[valueIndex], "%d", &intValue); err != nil {
					return fmt.Errorf("error parsing %s: %v", field, err)
				}
				switch field {
				case "ID":
					record.ID = intValue
				case "Age":
					record.Age = intValue
				}
			default:
				record.Department = row[valueIndex]
				switch field {
				case "Name":
					record.Name = row[valueIndex]
				case "Email":
					record.Email = row[valueIndex]
				}
			}
		}

		// Validate the record
		err = ValidateRecord(record)
		if err != nil {
			return fmt.Errorf("record validation failed: %v", err)
		}

		// Add the validated record to the list
		records = append(records, record)
	}

	// Write all records to the JSON file
	if err := jsonEncoder.Encode(records); err != nil {
		return fmt.Errorf("error writing JSON file: %v", err)
	}

	return nil
}

func main() {
	csvPath := "input.csv"
	jsonPath := "output.json"

	err := MigrateCSVToJSON(csvPath, jsonPath)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Migration successful. Data written to", jsonPath)
}
