package main

import (
	"fmt"
	"reflect"
)

type Employee1 struct {
	Name       string
	Age        int
	Department string
	Salary     float64
}

type Employee2 struct {
	Name     string
	Age      int
	Position string
	Salary   float64
	Bonus    float64
}

func main() {
	emp1 := []Employee1{
		{Name: "Alice", Age: 30, Department: "Engineering", Salary: 75000.0},
		{Name: "Bob", Age: 35, Department: "Sales", Salary: 60000.0},
		{Name: "Charlie", Age: 40, Department: "Operations", Salary: 50000.0},
	}

	emp2 := []Employee2{
		{Name: "Alice", Age: 30, Position: "Senior Engineer", Salary: 75000.0, Bonus: 5000.0},
		{Name: "Bob", Age: 36, Position: "Sales Manager", Salary: 62000.0, Bonus: 3000.0},
		{Name: "David", Age: 32, Position: "Marketing Specialist", Salary: 68000.0, Bonus: 4000.0},
	}

	compareSlices(emp1, emp2)
}

func compareSlices(slice1, slice2 interface{}) {
	v1 := reflect.ValueOf(slice1)
	v2 := reflect.ValueOf(slice2)

	if v1.Kind() != reflect.Slice || v2.Kind() != reflect.Slice {
		panic("both arguments must be slices")
	}

	// Create a map for quick lookup from slice1
	slice1Map := make(map[string]interface{})
	for i := 0; i < v1.Len(); i++ {
		elem := v1.Index(i)
		nameField := elem.FieldByName("Name")
		if nameField.IsValid() {
			name := nameField.String()
			slice1Map[name] = elem
		}
	}

	// Compare each element in slice2
	for i := 0; i < v2.Len(); i++ {
		elem2 := v2.Index(i)
		nameField := elem2.FieldByName("Name")
		if nameField.IsValid() {
			name := nameField.String()

			elem1, ok := slice1Map[name]
			if ok {
				compareStructs(elem1, elem2)
			} else {
				fmt.Printf("Employee %s not found in slice1.\n", name)
			}
		}
	}
}

func compareStructs(elem1, elem2 interface{}) {
	v1 := reflect.ValueOf(elem1)
	v2 := reflect.ValueOf(elem2)

	if v1.Kind() != reflect.Struct || v2.Kind() != reflect.Struct {
		panic("both arguments must be structs")
	}

	t1 := v1.Type()
	t2 := v2.Type()

	for i := 0; i < t1.NumField(); i++ {
		field1 := t1.Field(i)
		field2 := t2.Field(i)

		// Check if fields exist in both structs
		if field1.Name != field2.Name {
			continue
		}

		val1 := v1.Field(i)
		val2 := v2.Field(i)

		// Compare values
		if val1.Interface() != val2.Interface() {
			fmt.Printf("Difference in %s: %v vs %v\n", field1.Name, val1.Interface(), val2.Interface())
		}
	}

	// Check for extra fields in slice2
	for i := 0; i < t2.NumField(); i++ {
		field2 := t2.Field(i)
		if field1.Name != field2.Name {
			fmt.Printf("Extra field in slice2: %s\n", field2.Name)
		}
	}
}
