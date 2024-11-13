package main

import (
	"fmt"
	"reflect"
)

type StructWithFields struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func generateJSONMarshal(t reflect.Type) string {
	var code string
	code += "func (s " + t.Name() + ") MarshalJSON() ([]byte, error) {\n"
	code += "    result := make(map[string]interface{})\n"

	// Loop through the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")

		if tag != "" {
			code += fmt.Sprintf("    result[\"%s\"] = s.%s\n", tag, field.Name)
		}
	}

	code += "    bytes, err := json.Marshal(result)\n"
	code += "    return bytes, err\n"
	code += "}\n"

	return code
}

func main() {
	t := reflect.TypeOf(StructWithFields{})
	generatedCode := generateJSONMarshal(t)

	// Print the generated MarshalJSON method
	fmt.Println(generatedCode)
}
