package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age,omitempty"`
	Profile string `json:"profile,omitempty"`
}

// Custom MarshalJSON method to control the serialization logic
func (p Person) MarshalJSON() ([]byte, error) {
	// Use reflection to create a map of fields
	out := make(map[string]interface{})

	// Iterate over each field in the Person struct
	v := reflect.ValueOf(p)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// If the field's type is a reflect.Int, we will handle it specially
		if field.Type == reflect.TypeOf(1) {
			switch field.Name {
			case "Age":
				if v.Field(i).Int() > 18 {
					// Convert the age to a string
					out[field.Name] = strconv.Itoa(int(v.Field(i).Int()))
				}
			default:
				// For other int fields, just encode the value
				out[field.Name] = v.Field(i).Int()
			}
		} else {
			// For other field types, we directly add them to the map
			out[field.Name] = v.Field(i).Interface()
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
		Age:     16,
		Profile: "A user under 18",
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
