// risk.go

package main

func calculateRiskScore(age int, income float64, debt float64) float64 {
	// Simple risk calculation logic for demonstration
	riskScore := 0.1
	if age < 25 {
		riskScore += 0.2
	}
	if income < 50000 {
		riskScore += 0.3
	}
	if debt/income > 0.3 {
		riskScore += 0.4
	}
	return riskScore
}

func main() {
	// Main function not tested in this example
}
