package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

// LargeStruct represents a complex structure with potential for optimization
type LargeStruct struct {
	ID      int      `json:"id" xml:"id"`
	Name    string   `json:"name,omitempty" xml:"name,omitempty"`
	Items   []string `json:"items,omitempty" xml:"items>item,omitempty"`
	Details *Details `json:"details,omitempty" xml:"details,omitempty"`
}

// Details is a nested struct that might be large
type Details struct {
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	Quantity    int    `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

func serializeToJSON(data *LargeStruct) ([]byte, error) {
	return json.Marshal(data)
}

func serializeToXML(data *LargeStruct) ([]byte, error) {
	return xml.Marshal(data)
}

func main() {
	data := &LargeStruct{
		ID:   123,
		Name: "Sample Data",
		Items: []string{
			"Item1", "Item2", "Item3", "Item4", "Item5",
		},
		Details: &Details{
			Description: "Detailed Description",
			Quantity:    42,
		},
	}

	// JSON serialization
	jsonBytes, err := serializeToJSON(data)
	if err != nil {
		log.Fatalf("Error serializing to JSON: %v", err)
	}
	fmt.Println("Serialized JSON:")
	fmt.Println(string(jsonBytes))

	// XML serialization
	xmlBytes, err := serializeToXML(data)
	if err != nil {
		log.Fatalf("Error serializing to XML: %v", err)
	}
	fmt.Println("\nSerialized XML:")
	fmt.Println(string(xmlBytes))

	// Comparison of sizes
	log.Printf("\nSize of JSON: %d bytes", len(jsonBytes))
	log.Printf("Size of XML: %d bytes", len(xmlBytes))
}
