package main

import "fmt"

// calculateAverage calculates the average of a list of numbers.
// It returns 0 if the input array is empty.
func calculateAverage(numbers []float64) float64 {
	total := 0.0 // Use a descriptive variable name

	// Check if the input array is empty
	if len(numbers) == 0 {
		return 0
	}

	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

func main() {
	myList1 := []float64{1.0, 2.0, 3.0, 4.0}
	avg1 := calculateAverage(myList1)
	fmt.Println("Average of the list is:", avg1)

	myList2 := []float64{} // Empty list
	avg2 := calculateAverage(myList2)
	fmt.Println("Average of the empty list is:", avg2) // Output will be 0
}
