package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// CulturalHeritage struct represents a single cultural heritage dataset
type CulturalHeritage struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Country     string `json:"country"`
}

func main() {
	// Define the filename of the JSON file
	filename := "cultural_heritage.json"

	// Read the JSON file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		// Check for specific file-related errors
		switch {
		case os.IsNotExist(err):
			log.Fatalf("File %s not found: %v", filename, err)
		case os.IsPermission(err):
			log.Fatalf("Permission denied to read file %s: %v", filename, err)
		default:
			log.Fatalf("Error reading file %s: %v", filename, err)
		}
	}

	var culturalHeritages []CulturalHeritage

	// Unmarshal the JSON data into a slice of CulturalHeritage structs
	if err := json.Unmarshal(data, &culturalHeritages); err != nil {
		// Check for common JSON unmarshalling errors
		switch err.(type) {
		case *json.UnmarshalTypeError:
			log.Fatalf("Invalid JSON data type: %v", err)
		case *json.SyntaxError:
			log.Fatalf("Invalid JSON syntax: %v", err)
		default:
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
	}

	// Process each cultural heritage dataset using a range loop
	for _, heritage := range culturalHeritages {
		fmt.Printf("Name: %s\nDescription: %s\nCountry: %s\n\n", heritage.Name, heritage.Description, heritage.Country)
	}
}
