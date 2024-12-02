package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age,omitempty"`
	Profile string `json:"profile,omitempty"`
}

// Custom MarshalJSON method to control the serialization logic
func (p Person) MarshalJSON() ([]byte, error) {
	// Create a map to hold the fields we want to serialize
	out := make(map[string]interface{})

	// Iterate over each field in the Person struct
	v := reflect.ValueOf(p)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get the field's value
		fieldValue := v.Field(i)

		// Check if the field should be included in the JSON output
		if field.Name != "Age" || fieldValue.Int() != 0 {
			// Use the JSON tag if available, otherwise use the field name
			jsonTag := field.Tag.Get("json")
			if jsonTag != "" {
				out[jsonTag] = fieldValue.Interface()
			} else {
				out[field.Name] = fieldValue.Interface()
			}
		}
	}

	// Use the json package to encode the map to JSON
	return json.Marshal(out)
}

func main() {
	person1 := Person{
		Name:    "Alice",
		Age:     19,
		Profile: "A user over 18",
	}

	person2 := Person{
		Name:    "Bob",
		Age:     0,
		Profile: "A user with no age",
	}

	// Serialize each person to JSON
	json1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error serializing person1:", err)
		return
	}

	json2, err := json.Marshal(person2)
	if err != nil {
		fmt.Println("Error serializing person2:", err)
		return
	}

	// Print the serialized JSON
	fmt.Println(string(json1))
	fmt.Println(string(json2))
}
