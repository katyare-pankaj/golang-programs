package main

import (
	"fmt"
)

func calculateBMI(height float64, weight float64) string {
	bmi := weight / (height * height)

	if bmi < 18.5 {
		return "Underweight"
	} else if bmi < 25 {
		return "Normal"
	} else if bmi < 30 {
		return "Overweight"
	} else {
		return "Obese"
	}
}

func main() {
	height := 1.78
	weight := 75.5
	result := calculateBMI(height, weight)
	fmt.Println("BMI:", result)
}
