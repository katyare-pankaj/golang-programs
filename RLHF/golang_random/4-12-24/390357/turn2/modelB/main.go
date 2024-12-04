package main

import "fmt"

// Unused function, should be removed
func UnusedFunc() {
	// does nothing
}

// Used function
func UsedFunc() {
	fmt.Println("This function is used.")
}

// Properly documented function
// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

func ComplexFunction(cond1, cond2 bool) string {
	if cond1 {
		if cond2 {
			return "both true"
		} else {
			return "only cond1 true"
		}
	} else if cond2 {
		return "only cond2 true"
	}
	return "neither true"
}

func InefficientAppend() {
	var list []int
	for i := 0; i < 10; i++ {
		list = append(list, i)
	}
}

func EfficientAppend() {
	list := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		list = append(list, i)
	}
}

func main() {

	UsedFunc()
	add := Add(10, 20)
	fmt.Println("Add : ", add)

	str := ComplexFunction(true, false)
	fmt.Println("Result  : ", str)
}
