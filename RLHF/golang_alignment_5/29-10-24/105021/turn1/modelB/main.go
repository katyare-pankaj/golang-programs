package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gonum/stat"
)

func main() {
	// Generate random data
	rand.Seed(time.Now().UnixNano())
	n := 100
	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = rand.Float64() * 100
		y[i] = 2*x[i] + 5 + rand.NormFloat64()*10
	}

	// Fit linear regression model
	m, b := fitLinearRegression(x, y)

	// Predict resource allocation for a new value of x
	newX := 60.0
	predY := m*newX + b

	fmt.Println("Linear Regression Model:")
	fmt.Printf("Slope (m): %.2f\n", m)
	fmt.Printf("Intercept (b): %.2f\n", b)
	fmt.Printf("Predicted resource allocation for x=%.2f: %.2f\n", newX, predY)
}

// fitLinearRegression fits a linear regression model to the given data.
func fitLinearRegression(x, y []float64) (m, b float64) {
	// Calculate mean of x and y
	meanX := stat.Mean(x, nil)
	meanY := stat.Mean(y, nil)

	// Calculate covariance and variance of x and y
	cov := stat.Covariance(x, y, nil)
	varX := stat.Variance(x, nil)

	// Calculate slope (m) and intercept (b)
	m = cov / varX
	b = meanY - m*meanX

	return
}
