package main

import "fmt"

func calculateArea(shape string, length float64, width float64) float64 {
	area := 0.0
	switch shape {
	case "square":
		area = length * width
	case "rectangle":
		area = length * width
	case "triangle":
		area = 0.5 * length * width
	case "circle":
		area = 3.14 * length * length
	default:
		fmt.Println("Invalid shape!")
	}
	return area
}

func main() {
	result := calculateArea("triangle", 3.0, 4.0)
	fmt.Println("Area:", result)
}
