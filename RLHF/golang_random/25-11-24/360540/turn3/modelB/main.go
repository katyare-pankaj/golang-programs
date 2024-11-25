package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	// Read JSON schema from URL
	schemaURL := "https://microsoftedge.github.io/Demos/json-dummy-data/64KB.json"
	resp, err := http.Get(schemaURL)
	if err != nil {
		log.Fatalf("Error fetching schema: %v", err)
	}
	defer resp.Body.Close()
	schemaData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading schema response: %v", err)
	}

	// Read JSON data from a file
	dataFile := "data.json"
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		log.Fatalf("Error reading data file: %v", err)
	}

	// Create a JSON schema loader
	schemaLoader := gojsonschema.NewStringLoader(string(schemaData))

	// Create a JSON data loader
	dataLoader := gojsonschema.NewStringLoader(string(data))

	// Validate the data against the schema
	result, err := gojsonschema.Validate(schemaLoader, dataLoader)
	if err != nil {
		log.Fatalf("Error validating data: %v", err)
	}

	// Check if the validation was successful
	if !result.Valid() {
		fmt.Printf("Validation failed:\n")
		for _, err := range result.Errors() {
			fmt.Printf("- %s\n", err)
		}
	} else {
		fmt.Println("Validation successful!")
	}
}
