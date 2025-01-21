package main

import (
	"fmt"
)

func main() {
	// Define a dataset with some missing values represented by -999
	data := []float64{1.2, 2.5, -999, 3.7, 4.8, -999, 5.9, 6.1}

	// Calculate the mean of the valid values
	var sum float64
	var count int

	for _, value := range data {
		if value != -999 { // Check for missing value indicator
			sum += value
			count++
		}
	}

	// Ensure there is no division by zero
	if count == 0 {
		fmt.Println("No valid data to calculate mean.")
		return
	}

	mean := sum / float64(count)
	fmt.Printf("Mean of valid values: %.2f\n", mean)

	// Replace missing values with the calculated mean
	for i, value := range data {
		if value == -999 {
			data[i] = mean
		}
	}

	// Print the resultant dataset after imputation
	fmt.Println("Data after replacing missing values with mean:")
	fmt.Println(data)
}
