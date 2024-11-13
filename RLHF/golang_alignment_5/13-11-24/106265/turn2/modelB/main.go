package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// implement JSONMarshaler using reflection
func reflectJSONMarshal(v interface{}) ([]byte, error) {
	rt := reflect.ValueOf(v)
	if rt.Kind() != reflect.Struct {
		return nil, fmt.Errorf("cannot marshal non-struct type: %T", v)
	}

	m := make(map[string]interface{})
	rt = rt.Elem()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" {
			continue
		}

		m[tag] = field.Interface()
	}

	return json.Marshal(m)
}

func main() {
	type Point struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	point := &Point{
		X: 1,
		Y: 2,
	}

	// Generate and use the custom reflectJSONMarshal
	b, err := reflectJSONMarshal(point)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b)) // Output: {"x":1,"y":2}
}
