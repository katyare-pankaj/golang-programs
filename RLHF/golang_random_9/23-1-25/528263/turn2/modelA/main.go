package main

import (
	"fmt"
	"reflect"
)

// Map applies the provided function on each element of the input slice
// and returns a new slice containing the results.
func Map(slice interface{}, fn interface{}) interface{} {
	// Get the reflect.Value and reflect.Type of the input slice
	sliceVal := reflect.ValueOf(slice)
	sliceType := sliceVal.Type()

	// Ensure the first argument is a slice
	if sliceType.Kind() != reflect.Slice {
		panic("First argument must be a slice")
	}

	// Get the reflect.Value and reflect.Type of the function
	fnVal := reflect.ValueOf(fn)
	fnType := fnVal.Type()

	// Ensure function is valid with one input and one output
	if fnType.Kind() != reflect.Func || fnType.NumIn() != 1 || fnType.NumOut() != 1 {
		panic("Function must have one input and one output")
	}

	// Ensure function input type matches slice element type
	elemType := sliceType.Elem()
	if !elemType.AssignableTo(fnType.In(0)) {
		panic("Function input type must match slice element type")
	}

	// Create a new slice of the function's output type
	resultSliceType := reflect.SliceOf(fnType.Out(0))
	resultSlice := reflect.MakeSlice(resultSliceType, sliceVal.Len(), sliceVal.Len())

	// Apply the function to each element in the slice
	for i := 0; i < sliceVal.Len(); i++ {
		elem := sliceVal.Index(i)
		result := fnVal.Call([]reflect.Value{elem})
		resultSlice.Index(i).Set(result[0])
	}

	return resultSlice.Interface()
}

func main() {
	// Example usage with int slice and transformation function
	intSlice := []int{1, 2, 3, 4}
	mappedIntSlice := Map(intSlice, func(x int) int {
		return x * 2
	}).([]int)
	fmt.Println(mappedIntSlice) // Output: [2, 4, 6, 8]

	// Example usage with string slice and transformation function
	stringSlice := []string{"a", "b", "c"}
	mappedStringSlice := Map(stringSlice, func(s string) string {
		return s + s
	}).([]string)
	fmt.Println(mappedStringSlice) // Output: ["aa", "bb", "cc"]
}
