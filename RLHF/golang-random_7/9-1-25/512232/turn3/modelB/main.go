package main

import (
	"reflect"
	"unsafe"
)

const (
	// Size of a word on the target architecture (e.g., 32 or 64 bits)
	wordSize = 4 << (^uintptr(0) >> 63)
)

// securelyClearSlice clears sensitive data from large slices.
func securelyClearSlice(slice interface{}) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("Invalid argument type. Expecting a slice.")
	}

	// Get the slice's underlying array pointer and length
	ptr := unsafe.Pointer(s.Pointer())
	length := s.Len()

	// Check for empty slice
	if length == 0 {
		return
	}

	// Type switch on the slice's element type
	switch s.Type().Elem().Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		// For unsigned integer types, use word-sized operations for faster clearing
		words := length / wordSize
		remaining := length % wordSize

		for i := 0; i < words; i++ {
			*(*uintptr)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*wordSize)) = 0
		}

		// Clear the remaining bytes
		for i := words * wordSize; i < length; i++ {
			*(*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(i))) = 0
		}
	case reflect.String:
		// For string slices, zero out the string pointers
		for i := 0; i < length; i++ {
			*(*string)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*wordSize)) = ""
		}
	default:
		// For other types, use reflection to zero out each element
		for i := 0; i < length; i++ {
			elem := s.Index(i)
			elem.Set(reflect.Zero(elem.Type()))
		}
	}
}
