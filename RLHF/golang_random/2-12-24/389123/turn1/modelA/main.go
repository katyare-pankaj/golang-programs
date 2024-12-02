package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	ID       int    `json:"user_id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email"`
	Age      int    `json:"-"` // This field will be omitted during serialization
	IsActive bool   `json:"is_active"`
}

func customSerialize(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	// Get the reflect value of the interface
	rv := reflect.ValueOf(v)

	// Check if the value is a pointer
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	// Iterate through each field of the struct
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		typeInfo := rv.Type().Field(i)

		// Check if the field is marked with the json tag
		if tag, ok := typeInfo.Tag.Get("json"); ok {
			// Handle the "omitempty" directive
			if tag == "-" {
				continue
			}

			// Handle the struct field value
			if err := encoder.Encode(field.Interface()); err != nil {
				return nil, err
			}
		}
	}

	return buf.Bytes(), nil
}

func main() {
	user := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Age:      30,
		IsActive: true,
	}

	serializedData, err := customSerialize(&user)
	if err != nil {
		fmt.Println("Error serializing data:", err)
		return
	}

	fmt.Println("Serialized data:", string(serializedData))
}
