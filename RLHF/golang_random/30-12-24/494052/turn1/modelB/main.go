package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// CulturalHeritage struct represents a single cultural heritage dataset
type CulturalHeritage struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Country     string `json:"country"`
}

func main() {
	// Read the JSON data from a file
	data, err := ioutil.ReadFile("cultural_heritage.json")
	if err != nil {
		log.Fatal(err) // Fatal error, exit the program
	}

	var culturalHeritages []CulturalHeritage
	// Unmarshal the JSON data into a slice of CulturalHeritage structs
	if err := json.Unmarshal(data, &culturalHeritages); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Process each cultural heritage dataset using a range loop
	for _, heritage := range culturalHeritages {
		fmt.Printf("Name: %s\nDescription: %s\nCountry: %s\n\n", heritage.Name, heritage.Description, heritage.Country)
	}
}
