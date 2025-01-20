package main

import (
	"fmt"
	"reflect"
)

// Function to generate quiz questions from a slice using reflection
func generateQuizQuestions(slice interface{}) {
	val := reflect.ValueOf(slice)

	// Check if the provided interface is a slice
	if val.Kind() != reflect.Slice {
		fmt.Println("Provided data is not a slice.")
		return
	}

	// Iterate over the slice elements
	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		elemType := elem.Kind()

		// Generate questions based on the element type
		switch elemType {
		case reflect.String:
			fmt.Printf("Question %d: What is the value of this string? [%v]\n", i+1, elem.String())
		case reflect.Int:
			fmt.Printf("Question %d: What is %v plus 10?\n", i+1, elem.Int())
		case reflect.Float64:
			fmt.Printf("Question %d: Can you round %v to the nearest whole number?\n", i+1, elem.Float())
		default:
			fmt.Printf("Question %d: Unsupported data type %v\n", i+1, elemType)
		}
	}
}

func main() {
	stringSlice := []string{"apple", "banana", "cherry"}
	intSlice := []int{10, 20, 30}
	mixedSlice := []interface{}{"hello", 42, 3.14, true}

	fmt.Println("String Slice Questions:")
	generateQuizQuestions(stringSlice)

	fmt.Println("\nInt Slice Questions:")
	generateQuizQuestions(intSlice)

	fmt.Println("\nMixed Slice Questions:")
	generateQuizQuestions(mixedSlice)
}
