package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	// Read JSON data from a file
	dataFile := "data.json"
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		log.Fatalf("Error reading data file: %v", err)
	}

	// Read JSON schema from a file
	schemaFile := "schema.json"
	schema, err := ioutil.ReadFile(schemaFile)
	if err != nil {
		log.Fatalf("Error reading schema file: %v", err)
	}

	// Create a JSON schema loader
	schemaLoader := gojsonschema.NewStringLoader(string(schema))

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
