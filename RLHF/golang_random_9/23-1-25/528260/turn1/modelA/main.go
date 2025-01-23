package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

// SampleData represents a complex data structure
type SampleData struct {
	ID      int      `json:"id" xml:"id"`
	Name    string   `json:"name" xml:"name"`
	Items   []string `json:"items" xml:"items>item"`
	Details struct {
		Description string `json:"description" xml:"description"`
		Quantity    int    `json:"quantity" xml:"quantity"`
	} `json:"details" xml:"details"`
}

func main() {
	// Initialize the sample data
	data := SampleData{
		ID:    1,
		Name:  "Test Data",
		Items: []string{"Item1", "Item2", "Item3"},
		Details: struct {
			Description string `json:"description" xml:"description"`
			Quantity    int    `json:"quantity" xml:"quantity"`
		}{
			Description: "Sample description",
			Quantity:    99,
		},
	}

	// Serialize to JSON
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error serializing to JSON: %v", err)
	}
	fmt.Println("JSON Serialization:")
	fmt.Println(string(jsonBytes))

	// Deserialize back from JSON
	var jsonData SampleData
	err = json.Unmarshal(jsonBytes, &jsonData)
	if err != nil {
		log.Fatalf("Error deserializing from JSON: %v", err)
	}
	fmt.Printf("Deserialized JSON data: %+v\n\n", jsonData)

	// Serialize to XML
	xmlBytes, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error serializing to XML: %v", err)
	}
	fmt.Println("XML Serialization:")
	fmt.Println(string(xmlBytes))

	// Deserialize back from XML
	var xmlData SampleData
	err = xml.Unmarshal(xmlBytes, &xmlData)
	if err != nil {
		log.Fatalf("Error deserializing from XML: %v", err)
	}
	fmt.Printf("Deserialized XML data: %+v\n", xmlData)
}
