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
	ID      int       `json:"id" xml:"id"`
	Name    string    `json:"name" xml:"name"`
	Items   []*string `json:"items,omitempty" xml:"items>item,omitempty"` // Using pointers for slices
	Details *Details  `json:"details,omitempty" xml:"details,omitempty"`  // Using pointer to struct
}

// Details is a nested struct
type Details struct {
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	Quantity    int    `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

func main() {
	// Initialize the sample data
	data := SampleData{
		ID:   1,
		Name: "Test Data with Long Name and Lots of Description",
		Details: &Details{
			Description: strings.Repeat("X", 10000), // Large description
			Quantity:    999999,
		},
	}

	// Set a large number of items with short values for slicing performance demonstration
	numItems := 10000
	data.Items = make([]*string, numItems)
	itemValue := "item"
	for i := range data.Items {
		data.Items[i] = &itemValue
	}

	// Marshalling optimizations
	json.HTMLEscape(nil, []byte("")) // Preallocate for JSON strings

	// JSON Serialization
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error serializing to JSON: %v", err)
	}
	fmt.Println("JSON Serialization (optimized):")
	fmt.Printf("Size of serialized JSON: %d bytes\n", len(jsonBytes))

	// XML Serialization
	xmlBytes, err := xml.Marshal(data)
	if err != nil {
		log.Fatalf("Error serializing to XML: %v", err)
	}
	fmt.Println("XML Serialization (optimized):")
	fmt.Printf("Size of serialized XML: %d bytes\n", len(xmlBytes))
}
