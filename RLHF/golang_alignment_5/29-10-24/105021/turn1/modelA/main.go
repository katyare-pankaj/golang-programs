package main

import (
	"fmt"

	"github.com/gonum/floats"
	"github.com/gonum/stat"
)

// SimpleLinearRegression represents a simple linear regression model.
type SimpleLinearRegression struct {
	a float64 // slope
	b float64 // intercept
}

// Train trains the linear regression model on the given data points.
func (model *SimpleLinearRegression) Train(xs, ys []float64) {
	if len(xs) != len(ys) {
		panic("xs and ys must have the same length")
	}

	n := float64(len(xs))

	// Calculate sum of x, y, x^2, and x*y
	sumX := floats.Sum(xs)
	sumY := floats.Sum(ys)
	sumXX := stat.SumPow(xs, 2)
	sumXY := stat.Dot(xs, ys)

	// Calculate slope and intercept
	model.a = (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	model.b = (sumY / n) - (model.a*sumX)/n
}

// Predict predicts the value of y for a given x using the trained model.
func (model *SimpleLinearRegression) Predict(x float64) float64 {
	return model.a*x + model.b
}

func main() {
	// Sample data points: (x, y)
	xs := []float64{1, 2, 3, 4, 5}
	ys := []float64{2, 5, 8, 11, 14}

	// Train the linear regression model
	model := &SimpleLinearRegression{}
	model.Train(xs, ys)

	// Predict for a new value of x
	xNew := 6.0
	yPred := model.Predict(xNew)

	fmt.Println("Predicted value of y for x =", xNew, ":", yPred)

	// Calculate R-squared for model evaluation
	rSquared := calculateRSquared(xs, ys, model)
	fmt.Println("R-squared:", rSquared)
}

// calculateRSquared calculates the R-squared value for the model's performance.
func calculateRSquared(xs, ys []float64, model *SimpleLinearRegression) float64 {
	// Calculate total sum of squares
	sumSquaresTotal := stat.Variance(ys, nil) * float64(len(ys))

	// Calculate residual sum of squares
	var residuals []float64
	for i, x := range xs {
		yPred := model.Predict(x)
		residuals = append(residuals, ys[i]-yPred)
	}
	sumSquaresResidual := stat.Variance(residuals, nil) * float64(len(residuals))

	// Calculate R-squared
	return 1 - (sumSquaresResidual / sumSquaresTotal)
}
