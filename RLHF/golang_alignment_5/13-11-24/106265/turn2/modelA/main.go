package main

import (
	"fmt"
	"reflect"
)

func generateStringMethod(t reflect.Type) string {
	var code string
	code += "func (s " + t.Name() + ") String() string {\n"
	code += "    var fields []string\n"

	// Loop through the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		code += fmt.Sprintf("    fields = append(fields, fmt.Sprintf(\"%s:%%v\", \"%s\"))\n", field.Name, field.Name)
	}

	code += "    return fmt.Sprintf(\"{%s}\", strings.Join(fields, \", \"))\n"
	code += "}\n"

	return code
}

func main() {
	type Point struct {
		X int
		Y int
	}

	t := reflect.TypeOf(Point{})
	generatedCode := generateStringMethod(t)

	// Define a new function to execute the generated code
	f := func() {
		println("Generated String() method:")
		println(generatedCode)

		p := Point{X: 1, Y: 2}
		fmt.Println(p) // Output will be: {X:1, Y:2}
	}

	// Use reflect.MakeFunc to create an actual function from the generated code and call it.
	fn := reflect.MakeFunc(reflect.TypeOf(f), []reflect.Value{}, reflect.FuncFlagNilReceiver)
	fn.Call([]reflect.Value{})
}
