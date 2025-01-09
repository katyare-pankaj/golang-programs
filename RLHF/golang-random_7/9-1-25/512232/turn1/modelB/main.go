package main

import (
	"fmt"
	"reflect"
)

// ZeroSliceMemory zeroes out the memory allocated for a slice.
func ZeroSliceMemory(slice interface{}) {
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
			ZeroSliceMemory(elem.Interface())
		default:
			elem.Set(reflect.Zero(elem.Type()))
		}
	}
}

func main() {
	// Example usage:
	mySlice := make([]int, 100000)

	// Perform some operations on the slice
	for i := 0; i < len(mySlice); i++ {
		mySlice[i] = i
	}

	// Now, we can zero out the slice memory
	ZeroSliceMemory(mySlice)

	// Print the slice length to verify that it has been zeroed out
	fmt.Println(len(mySlice)) // Output: 100000
}
