package main

import (
	"fmt"
	"math"
	"time"

	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
)

// LinearRegression performs linear regression on the given data points.
func LinearRegression(xs, ys []float64) (slope, intercept float64) {
	n := len(xs)

	// Sum of x and y
	sumX := 0.0
	sumY := 0.0
	for _, x := range xs {
		sumX += x
	}
	for _, y := range ys {
		sumY += y
	}

	// Sum of x^2 and x*y
	sumXX := 0.0
	sumXY := 0.0
	for i := 0; i < n; i++ {
		x := xs[i]
		y := ys[i]
		sumXX += x * x
		sumXY += x * y
	}

	// Calculate slope and intercept
	denom := n*sumXX - sumX*sumX
	slope = (n*sumXY - sumX*sumY) / denom
	intercept = (sumY / float64(n)) - slope*(sumX/float64(n))

	return
}

func main() {
	// Generate some sample data for trend prediction
	const dataSize = 1_000_000
	xs := make([]float64, dataSize)
	ys := make([]float64, dataSize)
	for i := 0; i < dataSize; i++ {
		xs[i] = float64(i)
		ys[i] = 2.0*float64(i) + 1.0 + math.Randn()*0.1
	}

	// Perform linear regression
	start := time.Now()
	slope, intercept := LinearRegression(xs, ys)
	duration := time.Since(start)
	fmt.Printf("Slope: %f, Intercept: %f\n", slope, intercept)
	fmt.Printf("Execution time: %s\n", duration)

	// Using Apache Arrow for optimized linear regression
	start = time.Now()
	optimizedLinearRegression(xs, ys)
	duration = time.Since(start)
	fmt.Printf("Execution time using Apache Arrow: %s\n", duration)
}

// optimizedLinearRegression uses Apache Arrow for optimized linear regression
func optimizedLinearRegression(xs, ys []float64) {
	mem := memory.NewGoAllocator()
	defer mem.Free()

	xArray := array.NewFloat64Data(mem, arrow.NewFloat64Builder(mem), xs)
	yArray := array.NewFloat64Data(mem, arrow.NewFloat64Builder(mem), ys)

	// Calculate slope and intercept using Apache Arrow's linear regression function
	// (Note: This is just an example, and you would need to use the appropriate library function)
	// For this example, we'll just use a simple loop
	sumX := xArray.Sum()
	sumY := yArray.Sum()
	sumXX := xArray.Square().Sum()
	sumXY := xArray.Mul(yArray).Sum()

	n := float64(len(xs))
	denom := n*sumXX - sumX*sumX
	slope := (n*sumXY - sumX*sumY) / denom
	intercept := (sumY / n) - slope*(sumX/n)

	fmt.Printf("Slope: %f, Intercept: %f\n", slope, intercept)
}
