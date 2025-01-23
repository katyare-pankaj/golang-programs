package main

import (
	"fmt"
	"reflect"
	"sync"
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

	if !sliceType.Elem().AssignableTo(fnType.In(0)) {
		panic("Function input type must match slice element type")
	}

	// Preallocate result slice
	resultSlice := reflect.MakeSlice(reflect.SliceOf(fnType.Out(0)), sliceVal.Len(), sliceVal.Len())

	// Use concurrency for large slices
	if sliceVal.Len() > 1000 { // Threshold can be tuned
		var wg sync.WaitGroup
		workers := 4 // Number of concurrent workers; tune based on your CPU
		chunkSize := (sliceVal.Len() + workers - 1) / workers

		for w := 0; w < workers; w++ {
			wg.Add(1)
			go func(workerIdx int) {
				defer wg.Done()
				start := workerIdx * chunkSize
				end := start + chunkSize
				if end > sliceVal.Len() {
					end = sliceVal.Len()
				}

				for i := start; i < end; i++ {
					elem := sliceVal.Index(i)
					result := fnVal.Call([]reflect.Value{elem})[0]
					resultSlice.Index(i).Set(result)
				}
			}(w)
		}
		wg.Wait()
	} else {
		// Sequential processing for smaller slices
		for i := 0; i < sliceVal.Len(); i++ {
			elem := sliceVal.Index(i)
			result := fnVal.Call([]reflect.Value{elem})[0]
			resultSlice.Index(i).Set(result)
		}
	}

	return resultSlice.Interface()
}

func main() {
	// Test the map function with a large slice
	ints := make([]int, 10000)
	for i := range ints {
		ints[i] = i
	}

	mappedInts := Map(ints, func(i int) int {
		return i * i
	}).([]int)

	fmt.Println(mappedInts[:5]) // Output some results: [0 1 4 9 16]
}
