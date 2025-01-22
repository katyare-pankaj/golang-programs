package main  
import (  
    "fmt"
)

type intFunc func(int) int

func (f intFunc) apply(x int) int {
    return f(x)
}

// filterArray applies a type-safe callback function to each element in an array
// and returns a new array containing only the elements for which the callback returns true.
func filterArray(arr []int, predicate intFunc) []int {
    var filtered []int
    for _, value := range arr {
        if predicate.apply(value) {
            filtered = append(filtered, value)
        }
    }
    return filtered
}

// multiplyByTwo multiplies an integer by 2 and returns the result.
func multiplyByTwo(num int) int {
    return num * 2
}

// isEven returns true if an integer is even, false otherwise.
func isEven(num int) bool {
    return num % 2 == 0
}

// isPositive returns true if an integer is positive, false otherwise.
func isPositive(num int) bool {
    return num > 0
}

func main() {
    numbers := []int{-2, -1, 0, 1, 2, 3, 4, 5}

    fmt.Println("Original Array:", numbers)

    // Filter and transform even numbers using type-safe callbacks
    evenNumbers := filterArray(numbers, isEven)
    evenMultipliedByTwo := map(evenNumbers, multiplyByTwo)
    fmt.Println("Even Numbers multiplied by 2:", evenMultipliedByTwo)

    // Filter and transform positive numbers using type-safe callbacks
    positiveNumbers := filterArray(numbers, isPositive)
    positiveMultipliedByTwo := map(positiveNumbers, multiplyByTwo)
    fmt.Println("Positive Numbers multiplied by 2:", positiveMultipliedByTwo)

    // Define a custom callback for multiplying by 3
    multiplyByThree := func(num int) int {
        return num * 3
    }
    transformedArray := map(numbers, multiplyByThree)
    fmt.Println("Numbers multiplied by 3:", transformedArray)
}

func map(arr []int, transform intFunc) []int {
    var transformed []int
    for _, value := range arr {
        transformed = append(transformed, transform.apply(value))
    }
    return transformed
}