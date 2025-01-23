package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

// A Simple struct containing slices
type Simple struct {
	Values  []int    `json:"values" xml:"values>value"`
	Strings []string `json:"strings" xml:"strings>string"`
	Points  []Point  `json:"points" xml:"points>point"`
}

// Point struct representing a point in a 2D space
type Point struct {
	X int `json:"x" xml:"x"`
	Y int `json:"y" xml:"y"`
}

func main() {
	// Initialize a Simple struct with data
	data := Simple{
		Values:  []int{1, 2, 3, 4, 5},
		Strings: []string{"apple", "banana", "cherry"},
		Points:  []Point{{X: 1, Y: 2}, {X: 3, Y: 4}, {X: 5, Y: 6}},
	}

	// Encode to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON:", string(jsonData))

	// Encode to XML
	xmlData, err := xml.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("XML:", string(xmlData))

	// Preallocate buffer for efficient JSON unmarshaling
	buf := make([]byte, 0, 1024)
	buf = append(buf, jsonData...)

	// Unmarshal JSON back to the original struct
	var dataCopy Simple
	if err := json.Unmarshal(buf, &dataCopy); err != nil {
		panic(err)
	}
	fmt.Println("JSON Unmarshaled:", dataCopy)

	// Unmarshal XML back to the original struct
	if err := xml.Unmarshal(xmlData, &dataCopy); err != nil {
		panic(err)
	}
	fmt.Println("XML Unmarshaled:", dataCopy)

	// Write to a file (for demonstration purposes)
	jsonFile, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	if _, err := jsonFile.Write(jsonData); err != nil {
		panic(err)
	}

	xmlFile, err := os.Create("data.xml")
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	if _, err := xmlFile.Write(xmlData); err != nil {
		panic(err)
	}
}
