package main

import "fmt"

func calculateAverage(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers)) // Potential error: Division by zero
}

func main() {
	numbers := []float64{1.0, 2.0, 3.0}
	avg := calculateAverage(numbers)
	fmt.Println("Average:", avg)
}
