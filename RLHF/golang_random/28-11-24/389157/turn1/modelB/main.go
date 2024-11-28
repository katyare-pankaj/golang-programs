package main

import (
	"fmt"
	"math"
)

func main() {

	var a float64 = 3.14

	// Calculating the area of a circle with radius a
	area := math.Pi * a * a

	fmt.Println("The area of the circle is:", area)
}
