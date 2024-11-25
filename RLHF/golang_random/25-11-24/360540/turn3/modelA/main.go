package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	// URL of the JSON schema
	schemaURL := "https://json-schema.org/draft-07/schema#"

	// Load JSON schema from the URL
	schema, err := loadSchema(schemaURL)
	if err != nil {
		log.Fatalf("Error loading schema: %v", err)
	}

	// Read JSON data from a file
	dataFile := "data.json"
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		log.Fatalf("Error reading data file: %v", err)
	}

	// Convert data to a JSON graph to use in validation
	dataDoc, err := gojsonschema.NewDocument(gojsonschema.JSONLoader{Reader: ioutil.NopCloser(bytes.NewReader(data))})
	if err != nil {
		log.Fatalf("Error creating data graph: %v", err)
	}

	// Validate the data against the schema
	result, err := schema.Validate(dataDoc)
	if err != nil {
		log.Fatalf("Error validating data: %v", err)
	}

	// Check if the validation was successful
	if !result.Valid() {
		fmt.Printf("Validation failed:\n")
		for _, err := range result.Errors() {
			fmt.Printf("- %s\n", err.Description())
		}
	} else {
		fmt.Println("Validation successful!")
	}
}

func loadSchema(url string) (*gojsonschema.Schema, error) {
	// Send an HTTP GET request to load the schema
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get schema from %s: %v", url, err)
	}
	defer res.Body.Close()

	// Read the schema from the response body
	schemaBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema from %s: %v", url, err)
	}

	// Create a loader for the schema
	schemaLoader := gojsonschema.NewStringLoader(string(schemaBytes))

	// Validate and return the schema
	schema, err := gojsonschema.LoadSchema(schemaLoader)
	if err != nil {
		return nil, fmt.Errorf("failed to load schema: %v", err)
	}

	return schema, nil
}
