package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

// Process func uses reflection to process string fields concurrently
func processDataConcurrently(data interface{}) {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// Get number of struct fields
	numFields := v.NumField()
	// Create a wait group to synchronize concurrent goroutines
	var wg sync.WaitGroup
	// Ensure we have at least one goroutine running
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Start a goroutine for each field
	for i := 0; i < numFields; i++ {
		f := v.Field(i)
		wg.Add(1)
		go func(f reflect.Value) {
			defer wg.Done()
			// Check if the field is of type string
			if f.Type().Kind() == reflect.String {
				// Process the string field concurrently
				modifiedValue := fmt.Sprintf("Processed: %s", f.String())
				f.SetString(modifiedValue)
			}
		}(f)
	}
	// Wait for all goroutines to finish
	wg.Wait()
}

// Define some structs with string fields
type Person struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
	Age   int    `json:"age"`
}

type Car struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func main() {
	// Create instances of the structs
	person := Person{"Alice", "hiking", 25}
	car := Car{"Toyota", "Camry", 2021}

	dataStructures := []interface{}{
		person,
		car,
	}

	// Process each data structure concurrently
	for _, data := range dataStructures {
		processDataConcurrently(&data)
		fmt.Println(data)
	}

	fmt.Println("All data structures processed.")
}
