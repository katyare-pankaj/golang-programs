package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// ConcurrentMap applies the function fn concurrently to each element of the slice and returns a new slice
func ConcurrentMap(slice interface{}, fn interface{}, numGoroutines int) interface{} {
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

	if !sliceType.Elem().AssignableTo(fnType.In(0)) {
		panic("Function input type must match slice element type")
	}

	if numGoroutines <= 0 {
		numGoroutines = runtime.NumCPU()
	}

	resultSlice := reflect.MakeSlice(reflect.SliceOf(fnType.Out(0)), sliceVal.Len(), sliceVal.Len())

	ch := make(chan struct{}, numGoroutines)

	for i := 0; i < sliceVal.Len(); i++ {
		go func(idx int) {
			elem := sliceVal.Index(idx)
			result := fnVal.Call([]reflect.Value{elem})[0]
			resultSlice.Index(idx).Set(result)
			ch <- struct{}{}
		}(i)
	}

	for i := 0; i < sliceVal.Len(); i++ {
		<-ch
	}

	return resultSlice.Interface()
}

func main() {
	const numElements = 10000000
	data := make([]int, numElements)
	for i := 0; i < numElements; i++ {
		data[i] = i
	}

	// Measure performance with sequential and concurrent map
	mappedDataSeq := Map(data, func(x int) int {
		return x * 2
	}).([]int)

	mappedDataConcurrent := ConcurrentMap(data, func(x int) int {
		return x * 2
	}, 0).([]int)

	// Verify the correctness of the results
	for i := 0; i < numElements; i++ {
		if mappedDataSeq[i] != mappedDataConcurrent[i] {
			panic("Result mismatch!")
		}
	}

	fmt.Println("Performance test complete. Results match.")
}
