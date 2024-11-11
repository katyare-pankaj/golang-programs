package main

import (
	"fmt"
	"math"
)

func processData(data []float64) float64 {
	var sum float64
	for _, value := range data {
		sum += math.Pow(value, 2)
	}
	return sum
}

func main() {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	result := processData(data)
	fmt.Println("Result:", result)
}
