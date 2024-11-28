package main

import "fmt"

// old.go
func calculateSum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func computeAverage(values []float64) float64 {
	total := 0.0
	for _, val := range values {
		total += val
	}
	return total / float64(len(values))
}

func main() {

	sum := calculateSum([]int{1, 3, 4, 6, 7})
	fmt.Println("Calculate sum of numbers :", sum)

	total := computeAverage([]float64{1.0, 3.0, 4.0, 6.0, 7.0})
	fmt.Println("Calculate averate of numbers :", total)
}
