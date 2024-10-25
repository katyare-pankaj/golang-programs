package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 0; i < 10; i++ {
		x := i % 2
		// Use the branch prediction methods as needed
		result1 := predictStaticBranch(x)
		result2 := predictProfileGuided(x)
		result3 := indirectBranchPrediction(x > 0, f1, f2, x)

		fmt.Println(result1)
		fmt.Println(result2)
		fmt.Println(result3)

	}
}

func indirectBranchPrediction(condition bool, f1 func(int) int, f2 func(int) int, x int) int {
	var predFunc func(int) int
	if condition {
		predFunc = f1
	} else {
		predFunc = f2
	}
	return predFunc(x)
}

func f1(x int) int {
	return x * 2
}

func f2(x int) int {
	return x - 1
}

var branchPredictor = make(map[uintptr]bool)

func predictProfileGuided(x int) int {
	pc := make([]uintptr, 1)
	runtime.Callers(1, pc)
	key := pc[0]

	prediction := branchPredictor[key]
	if x > 0 {
		if !prediction {
			branchPredictor[key] = true
		}
		return x * 2
	}
	if prediction {
		branchPredictor[key] = false
	}
	return x - 1
}

func predictStaticBranch(x int) int {
	if x > 0 {
		return x * 2
	}
	return x - 1
}
