// riskassessment_test.go

package main

import (
	"testing"
)

func TestAssessRiskHigh(t *testing.T) {
	severity := 7
	likelihood := 6
	expected := "High"
	actual := assessRisk(severity, likelihood)
	if actual != expected {
		t.Errorf("Risk assessment failed. Expected: %s, Actual: %s", expected, actual)
	}
}

// Add more test cases for different risk severity and likelihood levels
func TestAssessRiskLow(t *testing.T) {
	severity := 2
	likelihood := 2
	expected := "Low"
	actual := assessRisk(severity, likelihood)
	if actual != expected {
		t.Errorf("Risk assessment failed. Expected: %s, Actual: %s", expected, actual)
	}
}
