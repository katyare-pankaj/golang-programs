package main

import (
	"fmt"
	"reflect"
)

type Employee1 struct {
	Name    string
	Age     int
	Salary  float64
	Address string
}

type Employee2 struct {
	Name       string
	Age        int
	Department string
	Salary     float64
}

func main() {
	emp1 := []Employee1{
		{Name: "Alice", Age: 30, Salary: 75000.0, Address: "New York"},
		{Name: "Bob", Age: 35, Salary: 60000.0, Address: "San Francisco"},
		{Name: "Charlie", Age: 40, Salary: 50000.0, Address: "London"},
	}

	emp2 := []Employee2{
		{Name: "Alice", Age: 30, Department: "Engineering", Salary: 75000.0},
		{Name: "Bob", Age: 36, Department: "Sales", Salary: 62000.0},
		{Name: "David", Age: 32, Department: "Marketing", Salary: 68000.0},
	}

	compareEmployees(emp1, emp2)
}

func compareEmployees(emp1 []Employee1, emp2 []Employee2) {
	fmt.Println("Comparing Employees:")
	fmt.Println("--------------------")

	emp1Map := make(map[string]Employee1)
	for _, emp := range emp1 {
		emp1Map[emp.Name] = emp
	}

	for _, emp2 := range emp2 {
		if emp1, ok := emp1Map[emp2.Name]; ok {
			compareStructs(emp1, emp2)
		} else {
			fmt.Printf("Employee %s not found in emp1.\n", emp2.Name)
		}
	}
}

func compareStructs(emp1 Employee1, emp2 Employee2) {
	fmt.Printf("\nComparing Employee %s:\n", emp1.Name)
	t1 := reflect.TypeOf(emp1)
	t2 := reflect.TypeOf(emp2)
	v1 := reflect.ValueOf(emp1)
	v2 := reflect.ValueOf(emp2)

	for i := 0; i < t1.NumField(); i++ {
		fieldName := t1.Field(i).Name
		field1 := v1.FieldByName(fieldName)

		if field2 := v2.FieldByName(fieldName); field2.IsValid() {
			if field1.Interface() != field2.Interface() {
				fmt.Printf("Difference in %s: %v vs %v\n", fieldName, field1.Interface(), field2.Interface())
			}
		} else {
			// Field not found in emp2
			fmt.Printf("Field %s not found in emp2.\n", fieldName)
		}
	}

	// Check for fields missing in emp1 but present in emp2
	for i := 0; i < t2.NumField(); i++ {
		fieldName := t2.Field(i).Name
		if _, found := t1.FieldByName(fieldName); !found {
			fmt.Printf("Field %s not found in emp1.\n", fieldName)
		}
	}
}
