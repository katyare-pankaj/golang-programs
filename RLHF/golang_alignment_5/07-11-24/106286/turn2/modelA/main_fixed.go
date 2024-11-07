package main

import "fmt"

func main() {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	result := make([]float64, len(data))

	// Unrolled loop for better performance
	for i := 0; i < len(data)-4; i += 5 {
		result[i] = data[i] * data[i]
		result[i+1] = data[i+1] * data[i+1]
		result[i+2] = data[i+2] * data[i+2]
		result[i+3] = data[i+3] * data[i+3]
		result[i+4] = data[i+4] * data[i+4]
	}

	// Remaining iterations for smaller arrays
	for i := 0; i < len(data); i++ {
		result[i] = data[i] * data[i]
	}

	fmt.Println("Result:", result)
}
