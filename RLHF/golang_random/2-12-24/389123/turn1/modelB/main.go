package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// Define our user struct
type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email,omitempty"`
	IsAdmin bool   `json:"is_admin,omitempty"`
}

// Custom MarshalJSON implementation to handle JSON serialization of User
func (u User) MarshalJSON() ([]byte, error) {
	var result strings.Builder
	result.WriteByte('{')
	v := reflect.ValueOf(u)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		// Skip fields with json:"omitempty" and value is zero value
		jsonTag := field.Tag.Get("json")
		if jsonTag == "omitempty" {
			switch value.(type) {
			case string:
				if value.(string) == "" {
					continue
				}
			case int:
				if value.(int) == 0 {
					continue
				}
			case bool:
				if !value.(bool) {
					continue
				}
			default:
				if reflect.ValueOf(value).IsNil() {
					continue
				}
			}
		}
		// Add field name and value to the result
		fieldName := jsonTag
		if fieldName == "" {
			fieldName = field.Name
		}
		if i > 0 {
			result.WriteByte(',')
		}
		result.WriteString(`"`)
		result.WriteString(fieldName)
		result.WriteString(`":`)
		// Encode the value based on its type
		valType := v.Field(i).Kind()
		switch valType {
		case reflect.String:
			result.WriteString(`"`)
			result.WriteString(fmt.Sprintf("%v", value))
			result.WriteString(`"`)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result.WriteString(fmt.Sprintf("%v", value))
		case reflect.Bool:
			result.WriteString(fmt.Sprintf("%v", value))
		default:
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
	// Create a User struct instance
	user := User{
		Name:    "Alice",
		Age:     25,
		Email:   "alice@example.com",
		IsAdmin: false,
	}
	// Serialize the user to JSON using custom MarshalJSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error serializing user:", err)
		return
	}
	// Print the serialized JSON
	fmt.Println(string(userJSON))
}
