package main

import (
	"fmt"
	"time"
)

type Thunk func() int

func AddThunk(a int, b int) Thunk {
	return func() int {
		return a + b
	}
}
func MulThunk(a int, b int) Thunk {
	return func() int {
		return a * b
	}
}
func main() {
	start := time.Now()
	addThunk := AddThunk(100000, 200000)
	mulThunk := MulThunk(addThunk(), 300000)

	// Computation is not performed until this point
	result := mulThunk()
	fmt.Println("Result:", result)
	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)
}
