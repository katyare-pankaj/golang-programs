package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bmatsuo/go-jsonschema"
)

// Define your JSON data structure
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Define the JSON schema as a string
	schemaJSON := `
	{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"type": "object",
		"properties": {
			"name": {
				"type": "string",
				"minLength": 1
			},
			"age": {
				"type": "integer",
				"minimum": 0
			},
			"email": {
				"type": "string",
				"format": "email"
			}
		},
		"required": ["name", "age", "email"]
	}
	`

	// Load the JSON data from a file
	dataFile := "data.json"
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		log.Fatalf("Error reading data file: %v", err)
	}

	// Parse the JSON data into a User struct
	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Fatalf("Error parsing JSON data: %v", err)
	}

	// Parse the JSON schema
	schemaLoader := jsonschema.NewLoader(jsonschema.URLFromString(""))
	schema, err := schemaLoader.Load("data:, " + schemaJSON)
	if err != nil {
		log.Fatalf("Error parsing JSON schema: %v", err)
	}

	// Validate the JSON data against the schema
	validator := jsonschema.NewValidator(schema)
	result, err := validator.ValidateBytes(data)
	if err != nil {
		log.Fatalf("Error validating JSON data: %v", err)
	}

	// Log any validation errors
	if !result.Valid() {
		for _, err := range result.Errors() {
			fmt.Printf("Validation error: %v\n", err)
		}
	} else {
		fmt.Println("JSON data is valid against the schema.")
	}
}
