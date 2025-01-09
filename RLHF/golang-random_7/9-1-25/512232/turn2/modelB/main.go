package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// securelyClearSlice clears sensitive data in a slice by setting each element to its zero value
func securelyClearSlice(slice interface{}) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("Invalid argument type. Expecting a slice.")
	}

	for i := 0; i < s.Len(); i++ {
		elem := s.Index(i)
		switch elem.Kind() {
		case reflect.String:
			elem.SetString("")
		case reflect.Slice, reflect.Array:
			securelyClearSlice(elem.Interface())
		default:
			elem.Set(reflect.Zero(elem.Type()))
		}
	}
}

func main() {
	// Example usage:
	sensitiveData := make([]byte, 1024)

	// Perform some operations on the sensitive data
	for i := 0; i < len(sensitiveData); i++ {
		sensitiveData[i] = 'S'
	}

	// Now, we can securely clear the slice
	securelyClearSlice(sensitiveData)

	// Run garbage collection to free up the memory
	runtime.GC()

	// Print the slice to verify that it has been zeroed out
	fmt.Println(sensitiveData) // Output: [0 0 0 ... 0 0]
}
