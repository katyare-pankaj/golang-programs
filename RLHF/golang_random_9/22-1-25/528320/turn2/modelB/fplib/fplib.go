package fplib

// Iterable defines the interface that any slice-like data structure can conform to,
// enabling it to be processed using the fplib functions.
type Iterable interface {
	Len() int
	Index(i int) interface{}
}

// Func is the base type for all callback functions in fplib.
// It represents a function that takes a single argument of type 'T' and returns a result of type 'U'.
type Func[T, U any] func(T) U

// Predicate is a type-safe callback function that takes an argument of type 'T' and returns 'bool'.
type Predicate[T any] func(T) bool

// Map applies a function to each element of an iterable and returns a new slice of results.
// The result type 'U' is inferred from the callback function 'fn'.
func Map[T, U any](iterable Iterable, fn Func[T, U]) []U {
	result := make([]U, iterable.Len())
	for i := 0; i < iterable.Len(); i++ {
		value := iterable.Index(i).(T)
		result[i] = fn(value)
	}
	return result
}

// Filter returns a new slice containing only the elements that satisfy a given predicate.
func Filter[T any](iterable Iterable, predicate Predicate[T]) []T {
	var result []T
	for i := 0; i < iterable.Len(); i++ {
		value := iterable.Index(i).(T)
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

// Reduce applies a function to each element of an iterable, accumulating the result.
func Reduce[T any](iterable Iterable, fn Func[T, T], initial T) T {
	accumulator := initial
	for i := 0; i < iterable.Len(); i++ {
		value := iterable.Index(i).(T)
		accumulator = fn(accumulator, value)
	}
	return accumulator
}

// SliceIterator is an implementation of the Iterable interface for slices.
type SliceIterator[T any] struct {
	slice []T
}

func (s SliceIterator[T]) Len() int {
	return len(s.slice)
}

func (s SliceIterator[T]) Index(i int) interface{} {
	return s.slice[i]
}

// FuncOf creates a new Func instance from a regular Go function.
// This function is convenient for wrapping existing functions into Func types for use in fplib functions.
func FuncOf[T, U any](f func(T) U) Func[T, U] {
	return f
}

// PredicateOf creates a new Predicate instance from a regular Go function.
func PredicateOf[T any](f func(T) bool) Predicate[T] {
	return f
}
