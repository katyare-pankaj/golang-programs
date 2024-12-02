package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// Person struct containing Name and Age fields
type Person struct {
	Name string `json:"include"`
	Age  int    `json:"include"`
	City string `json:"-"` // We'll exclude this field using "-" tag
}

// Custom MarshalJSON implementation
func (p Person) MarshalJSON() ([]byte, error) {
	var result strings.Builder
	result.WriteByte('{')
	v := reflect.ValueOf(p)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")

		// Skip fields with '-' tag or no 'json' tag
		if jsonTag == "-" || jsonTag == "" {
			continue
		}

		// Get the value of the current field
		value := v.Field(i).Interface()

		// Check if the field is tagged with 'include'
		if strings.Contains(jsonTag, "include") {
			// Field should be included in the JSON output
			fieldName := strings.Split(jsonTag, ",")[0] // Extract the field name from the tag
			if i > 0 {
				result.WriteByte(',')
			}
			result.WriteString(`"`)
			result.WriteString(fieldName)
			result.WriteString(`":`)
			// Encode the value based on its type
			valBytes, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			result.Write(valBytes)
		}
	}
	result.WriteByte('}')
	return []byte(result.String()), nil
}

func main() {
	person := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}
	// Serialize the person to JSON using the custom MarshalJSON implementation
	personJSON, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error serializing person:", err)
		return
	}
	// Print the serialized JSON
	fmt.Println(string(personJSON))
}
