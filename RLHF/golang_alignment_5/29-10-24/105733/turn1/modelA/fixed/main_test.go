package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// RiskAssessment represents a risk assessment
type RiskAssessment struct {
	ID          int
	Description string
	Score       int
}

// AssessRisk assesses a risk based on some criteria
func AssessRisk(assessment RiskAssessment) string {
	if assessment.Score >= 80 {
		return "High"
	} else if assessment.Score >= 60 {
		return "Medium"
	} else {
		return "Low"
	}
}

// String method to format RiskAssessment as a string
func (ra RiskAssessment) String() string {
	return fmt.Sprintf("ID: %d, Description: %s, Score: %d", ra.ID, ra.Description, ra.Score)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	assessment := RiskAssessment{
		ID:          1,
		Description: "Sample Risk Assessment",
		Score:       rand.Intn(100),
	}
	fmt.Printf("Risk Assessment: %s\n", assessment) // This will now work
	fmt.Printf("Risk Level: %s\n", AssessRisk(assessment))
}

// TestAssessRisk tests the AssessRisk function
func TestAssessRisk(t *testing.T) {
	tests := []struct {
		name     string
		score    int
		expected string
	}{
		{"High Risk", 80, "High"},
		{"Medium Risk", 60, "Medium"},
		{"Low Risk", 59, "Low"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AssessRisk(RiskAssessment{Score: tt.score})
			if result != tt.expected {
				t.Errorf("got %s, want %s", result, tt.expected)
			}
		})
	}
}
