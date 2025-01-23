package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

// SampleData represents a complex data structure
type SampleData struct {
	ID      int      `json:"id" xml:"id"`
	Name    string   `json:"name" xml:"name"`
	Items   []string `json:"items,omitempty" xml:"items>item,omitempty"`
	Details *Details `json:"details,omitempty" xml:"details,omitempty"`
}

// Details is a nested struct to keep data separated
type Details struct {
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	Quantity    int    `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

func main() {
	// Initialize the sample data
	data := SampleData{
		ID:   1,
		Name: "Test Data",
		Details: &Details{
			Description: "Sample description",
			Quantity:    99,
		},
	}
	// Set a large number of items for slicing performance demonstration
	data.Items = strings.Split("abcde fghij klmno pqrst uvwxy z1234 56789", " ")

	// JSON Serialization with optimization
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error serializing to JSON: %v", err)
	}
	fmt.Println("JSON Serialization (optimized):")
	fmt.Println(string(jsonBytes))

	// JSON Deserialization
	var jsonData SampleData
	if err = json.Unmarshal(jsonBytes, &jsonData); err != nil {
		log.Fatalf("Error deserializing from JSON: %v", err)
	}
	fmt.Printf("Deserialized JSON data: %+v\n\n", jsonData)

	// XML Serialization with optimization
	xmlBytes, err := xml.Marshal(data)
	if err != nil {
		log.Fatalf("Error serializing to XML: %v", err)
	}
	fmt.Println("XML Serialization (optimized):")
	fmt.Println(string(xmlBytes))

	// XML Deserialization
	var xmlData SampleData
	if err = xml.Unmarshal(xmlBytes, &xmlData); err != nil {
		log.Fatalf("Error deserializing from XML: %v", err)
	}
	fmt.Printf("Deserialized XML data: %+v\n", xmlData)
}
