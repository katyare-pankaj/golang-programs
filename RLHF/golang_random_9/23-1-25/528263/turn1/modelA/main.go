package main

import (
	"fmt"
	"reflect"
)

// Map applies the function fn to each element of the slice and returns a new slice
func Map(slice interface{}, fn interface{}) interface{} {
	// Get the Type of the input slice
	sliceVal := reflect.ValueOf(slice)
	sliceType := sliceVal.Type()

	// Ensure the first argument is a slice
	if sliceType.Kind() != reflect.Slice {
		panic("First argument must be a slice")
	}

	// Get the Type of the function
	fnVal := reflect.ValueOf(fn)
	fnType := fnVal.Type()

	// Ensure function has one input and one output
	if fnType.Kind() != reflect.Func || fnType.NumIn() != 1 || fnType.NumOut() != 1 {
		panic("Function must have exactly one input and one output parameter")
	}

	// Ensure function input type matches slice element type
	if !sliceType.Elem().AssignableTo(fnType.In(0)) {
		panic("Function input type must match slice element type")
	}

	// Create a new slice to hold the results
	resultSlice := reflect.MakeSlice(reflect.SliceOf(fnType.Out(0)), sliceVal.Len(), sliceVal.Len())

	// Apply the function to each element of the slice
	for i := 0; i < sliceVal.Len(); i++ {
		elem := sliceVal.Index(i)
		result := fnVal.Call([]reflect.Value{elem})
		resultSlice.Index(i).Set(result[0])
	}

	return resultSlice.Interface()
}

func main() {
	// Example usage
	slice := []int{1, 2, 3, 4}
	mappedSlice := Map(slice, func(x int) int {
		return x * 2
	}).([]int)

	fmt.Println(mappedSlice) // Output: [2, 4, 6, 8]
}
