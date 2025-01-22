package funclib

// Map applies a function to each element of a slice and returns a new slice of results.
func Map[T any, U any](slice []T, fn func(T) U) []U {
	mapped := make([]U, len(slice))
	for i, v := range slice {
		mapped[i] = fn(v)
	}
	return mapped
}

// Filter returns a new slice containing only the elements that satisfy the predicate function.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var filtered []T
	for _, v := range slice {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// Reduce applies a function to each element of a slice, accumulating the result.
func Reduce[T any, U any](slice []T, fn func(U, T) U, initial U) U {
	accumulator := initial
	for _, v := range slice {
		accumulator = fn(accumulator, v)
	}
	return accumulator
}
