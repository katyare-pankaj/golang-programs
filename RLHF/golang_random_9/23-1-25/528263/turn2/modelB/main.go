package main

import (
	"fmt"
	"reflect"
)

func Map(slice interface{}, fn interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	sliceType := sliceVal.Type()

	if sliceType.Kind() != reflect.Slice {
		panic("First argument must be a slice")
	}

	fnVal := reflect.ValueOf(fn)
	fnType := fnVal.Type()

	if fnType.Kind() != reflect.Func || fnType.NumIn() != 1 || fnType.NumOut() != 1 {
		panic("Function must have exactly one input and one output parameter")
	}

	// Ensure the function input type matches the slice element type
	if !sliceType.Elem().AssignableTo(fnType.In(0)) {
		panic("Function input type must match slice element type")
	}

	// Create a new slice with the result type of the function
	resultSlice := reflect.MakeSlice(reflect.SliceOf(fnType.Out(0)), sliceVal.Len(), sliceVal.Len())

	// Loop through the slice elements and apply the function
	for i := 0; i < sliceVal.Len(); i++ {
		elem := sliceVal.Index(i)
		result := fnVal.Call([]reflect.Value{elem})[0]
		resultSlice.Index(i).Set(result)
	}

	return resultSlice.Interface()
}

func main() {
	// Test the map function with different slice types
	strings := []string{"a", "b", "c"}
	ints := []int{1, 2, 3}
	floats := []float64{1.0, 2.0, 3.0}

	mappedStrings := Map(strings, func(s string) string {
		return s + "x"
	}).([]string)

	mappedInts := Map(ints, func(i int) int {
		return i * i
	}).([]int)

	mappedFloats := Map(floats, func(f float64) float64 {
		return f * 2
	}).([]float64)

	fmt.Println(mappedStrings) // Output: [ax bx cx]
	fmt.Println(mappedInts)    // Output: [1 4 9]
	fmt.Println(mappedFloats)  // Output: [2 4 6]
}
