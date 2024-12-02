package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age,omitempty"`
	Height float64 `json:"height,omitempty"`
	City   string  `json:"city,omitempty"`
}

func (p Person) MarshalJSON() ([]byte, error) {
	out := make(map[string]interface{})
	v := reflect.ValueOf(p)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")

		// Check if the field's value is zero
		isZero := false
		switch v.Field(i).Kind() {
		case reflect.String:
			isZero = v.Field(i).String() == ""
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			isZero = v.Field(i).Int() == 0
		case reflect.Float32, reflect.Float64:
			isZero = v.Field(i).Float() == 0
		case reflect.Bool:
			isZero = !v.Field(i).Bool()
		default:
			isZero = v.Field(i).IsNil()
		}

		// Exclude the field if its value is zero
		if !isZero {
			fieldName := jsonTag
			if fieldName == "" {
				fieldName = field.Name
			}
			out[fieldName] = v.Field(i).Interface()
		}
	}

	return json.Marshal(out)
}

func main() {
	person1 := Person{
		Name:   "Alice",
		Age:    18,
		Height: 170.5,
		City:   "New York",
	}

	person2 := Person{
		Name: "Bob",
		City: "Paris",
	}

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

	fmt.Println(string(json1))
	fmt.Println(string(json2))
}
