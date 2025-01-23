package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// ProcessFunc defines the type of function that transforms data in the pipeline
type ProcessFunc func([]int) []int

// transformStage is an example of a stage in the pipeline that processes data
func transformStage(data []int) []int {
	// Example transformation: Multiply each element by 2
	for i := range data {
		data[i] *= 2
	}
	return data
}

// unsafeAppend appends elements to a slice without allocating new memory
func unsafeAppend(s *[]int, elements []int) {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(s))
	elementHeader := (*reflect.SliceHeader)(unsafe.Pointer(&elements))

	// Ensure there is enough capacity
	if sliceHeader.Cap < sliceHeader.Len+elementHeader.Len {
		panic("not enough capacity in slice")
	}

	// Copy elements to the destination slice
	dest := unsafe.Pointer(uintptr(sliceHeader.Data) + uintptr(sliceHeader.Len)*unsafe.Sizeof((*s)[0]))
	src := unsafe.Pointer(elementHeader.Data)
	size := uintptr(elementHeader.Len) * unsafe.Sizeof((*elements)[0])

	unsafe.Copy(dest, src, size)
	sliceHeader.Len += elementHeader.Len
}

// runPipeline executes the data pipeline
func runPipeline(data []int, stages ...ProcessFunc) []int {
	for _, stage := range stages {
		data = stage(data)
	}
	return data
}

func main() {
	data := make([]int, 0, 10)
	initialData := []int{1, 2, 3, 4, 5}

	// Using unsafeAppend to add initial data without reallocating
	unsafeAppend(&data, initialData)

	// Define pipeline stages
	stages := []ProcessFunc{
		transformStage,
	}

	// Run data through the pipeline
	finalData := runPipeline(data, stages...)

	// Print results
	fmt.Println("Final Data:", finalData)
}
