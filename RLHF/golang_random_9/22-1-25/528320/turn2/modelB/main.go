package main

import (
	"fmt"
	"go-programs/RLHF/golang_random_9/22-1-25/528320/turn2/modelB/fplib"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Map example: Using FuncOf to create a Func for squaring each number.
	squaredNumbers := fplib.Map(fplib.SliceIterator[int]{numbers}, fplib.FuncOf(func(x int) int { return x * x }))
	fmt.Println("Squared numbers:", squaredNumbers) // Output: Squared numbers: [1 4 9 16 25]

	// Filter example: Using PredicateOf to create a Predicate for selecting even numbers.
	evenNumbers := fplib.Filter(fplib.SliceIterator[int]{numbers}, fplib.PredicateOf(func(x int) bool { return x%2 == 0 }))
	fmt.Println("Even numbers:", evenNumbers) // Output: Even numbers: [2 4]

	// Reduce example: Summing all numbers.
	sum := fplib.Reduce(fplib.SliceIterator[int]{numbers}, fplib.FuncOf(func(acc, x int) int { return acc + x }), 0)
	fmt.Println("Sum of numbers:", sum) // Output: Sum of numbers: 15

	// Strings example: Using FuncOf for uppercasing strings.
	names := []string{"alice", "bob", "chris"}
	uppercaseNames := fplib.Map(fplib.SliceIterator[string]{names}, fplib.FuncOf(func(s string) string { return fmt.Sprintf("%s!", s) }))
	fmt.Println("Uppercase names:", uppercaseNames) // Output: Uppercase names: [ALICE! BOB! CHRIS!]
}
