// risk_test.go

package main

import (
	"testing"
)

func TestCalculateRiskScore(t *testing.T) {
	type testCase struct {
		name     string
		age      int
		income   float64
		debt     float64
		expected float64
	}
	testCases := []testCase{
		{name: "Young High Income", age: 22, income: 70000, debt: 10000, expected: 0.6},
		{name: "Old Low Income", age: 65, income: 30000, debt: 8000, expected: 0.7},
		// Add more test cases
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := calculateRiskScore(tc.age, tc.income, tc.debt)
			if actual != tc.expected {
				t.Errorf("Calculated risk score (%f) doesn't match expected (%f)", actual, tc.expected)
			}
		})
	}
}
