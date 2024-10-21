package main

import "fmt"

func calculateAverage(numbers []float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

func main() {
	myList := []float64{1.0, 2.0, 3.0, 4.0}
	avg := calculateAverage(myList)
	fmt.Println("Average of the list is:", avg)
}
