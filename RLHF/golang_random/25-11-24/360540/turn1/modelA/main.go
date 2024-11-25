package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Schema represents the JSON schema structure
type Schema struct {
	Fields map[string]Field `json:"fields"`
}

// Field represents a field in the schema
type Field struct {
	Type    string      `json:"type"`
	Default interface{} `json:"default,omitempty"`
}

// Data represents the incoming JSON data
type Data struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// Validate validates incoming data against the schema
func Validate(schema *Schema, data interface{}) bool {
	// Implement validation logic here
	// For simplicity, we assume all fields are present and valid
	return true
}

// UpdateSchema updates the schema based on incoming data
func UpdateSchema(schema *Schema, data interface{}) {
	dataMap := data.(map[string]interface{})

	for key, value := range dataMap {
		if _, exists := schema.Fields[key]; !exists {
			var fieldType string
			var defaultValue interface{}

			switch v := value.(type) {
			case string:
				fieldType = "string"
				defaultValue = v
			case int:
				fieldType = "int"
				defaultValue = v
			case float64:
				fieldType = "float64"
				defaultValue = v
			default:
				// Unsupported types
				fmt.Println("Unsupported type:", v)
				continue
			}

			schema.Fields[key] = Field{Type: fieldType, Default: defaultValue}
		}
	}
}

func main() {
	// Load existing schema from a file
	schemaFile := "schema.json"
	schema := &Schema{Fields: make(map[string]Field)}

	data, err := ioutil.ReadFile(schemaFile)
	if err == nil {
		err = json.Unmarshal(data, schema)
		if err != nil {
			fmt.Println("Error unmarshaling schema:", err)
		}
	}

	// Example incoming data
	//dataIn := Data{Name: "Alice", Age: 30, Email: "alice@example.com"}
	dataMap := make(map[string]interface{})
	json.Unmarshal([]byte(`{"name": "Bob", "age": 25, "city": "New York"}`), &dataMap)

	// Validate the data against the schema
	if Validate(schema, dataMap) {
		fmt.Println("Data is valid.")
	} else {
		fmt.Println("Data is invalid.")
	}

	// Update the schema with new fields
	UpdateSchema(schema, dataMap)

	// Save the updated schema back to the file
	dataOut, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling schema:", err)
	}
	err = ioutil.WriteFile(schemaFile, dataOut, 0644)
	if err != nil {
		fmt.Println("Error writing schema to file:", err)
	}

	fmt.Println("Schema updated and saved.")
}
