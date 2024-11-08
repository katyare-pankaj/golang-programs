package main

import "fmt"

// Refactored Code (Functional Approach)

func calculateAreaFunctional(radius float64) float64 {
	// Pure function to calculate the square
	square := func(x float64) float64 {
		return x * x
	}

	// Pure function to calculate the area using pi
	areaWithPi := func(r float64) float64 {
		return 3.14 * square(r)
	}

	return areaWithPi(radius)
}

func main() {

	area := calculateAreaFunctional(5)

	fmt.Println("Area =", area)
}
