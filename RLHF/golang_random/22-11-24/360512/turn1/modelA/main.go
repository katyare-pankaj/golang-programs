package main

import (
	"fmt"
	"reflect"
	"sync"
)

// Process func uses reflection to process any data structure with a "process" method
func processData(data interface{}) {
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

		// For example, we increment or modify the string value
		modifiedValue := fmt.Sprintf("Processed: %s", f.String())
		f.SetString(modifiedValue)
	}
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

	// Create goroutines to process each data structure
	for _, data := range dataStructures {
		wg.Add(1)
		go func(d interface{}) {
			defer wg.Done()
			processData(d)
			fmt.Println(d)
		}(data)
	}

	wg.Wait()
	fmt.Println("All data structures processed.")
}
