package main

import (
	"fmt"
	"reflect"
	"sync"
)

// ProcessStructs function uses reflection to process all string fields in a struct concurrently
func processStructs(data interface{}) {
	var wg sync.WaitGroup
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		if !f.CanSet() {
			continue
		}

		if f.Type().Kind() != reflect.String {
			continue
		}

		wg.Add(1)
		go func(field reflect.Value) {
			defer wg.Done()
			// Alter the string field
			modifiedValue := fmt.Sprintf("Processed: %s", field.String())
			field.SetString(modifiedValue)
		}(f)
	}

	wg.Wait()
}

// Define the types of data we want to process
type Person struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}

type Car struct {
	Make  string `json:"make"`
	Model string `json:"model"`
}

func main() {
	var wg sync.WaitGroup

	dataStructures := []interface{}{
		Person{"Alice", "hiking"},
		Person{"Bob", "reading"},
		Car{"Toyota", "Camry"},
		Car{"Honda", "Civic"},
	}

	// Process each struct concurrently
	for _, data := range dataStructures {
		wg.Add(1)
		go func(d interface{}) {
			defer wg.Done()
			processStructs(d)
			fmt.Println(d)
		}(data)
	}

	wg.Wait()
	fmt.Println("All data structures processed.")
}
