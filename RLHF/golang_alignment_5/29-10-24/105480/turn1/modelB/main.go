

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/gcausal"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Generate random data for two variables: steps and sleep hours
	rand.Seed(time.Now().UnixNano())
	n := 100
	steps := make([]float64, n)
	sleepHours := make([]float64, n)
	for i := 0; i < n; i++ {
		steps[i] = rand.NormFloat64()*1000 + 5000
		sleepHours[i] = rand.NormFloat64()*7 + 8
	}

	// Assume there is a causal relationship between steps and sleep hours (steps = b0 + b1*sleepHours + e)
	b0 := 5000.0
	b1 := 100.0
	sigma := 500.0
	trueSteps := make([]float64, n)
	for i := 0; i < n; i++ {
		trueSteps[i] = b0 + b1*sleepHours[i] + rand.NormFloat64()*sigma
	}

	// Perform causal filtering using GCEA (Generalized Causal Estimation Algorithm)
	causalModel := gcausal.NewGaussianCOPT(trueSteps, sleepHours)
	causalModel.Fit()
	fittedSteps := causalModel.Predict(sleepHours)

	// Calculate the residuals (difference between measured and fitted steps)
	residuals := make([]float64, n)
	for i := 0; i < n; i++ {
		residuals[i] = steps[i] - fittedSteps[i]
	}

	// Filter out noisy data points based on a threshold (e.g., standard deviation of the residuals)
	threshold := 3 * gcausal.StdDev(residuals)
	filteredSteps := make([]float64, 0)
	filteredSleepHours := make([]float64, 0)
	for i := 0; i < n; i++ {
		if abs(residuals[i]) <= threshold {
			filteredSteps = append(filteredSteps, steps[i])
			filteredSleepHours = append(filteredSleepHours, sleepHours[i])
		}
	}

	// Plot the original, fitted, and filtered data
	p, err := plot.New()
