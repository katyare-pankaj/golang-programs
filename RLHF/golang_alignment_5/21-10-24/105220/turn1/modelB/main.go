package main

import (
	"fmt"

	"github.com/go-fuzzy/fuzzy"
)

// Define fuzzy sets for age and income
type FuzzySet struct {
	Name string
	Func *fuzzy.Func
}

func main() {
	// Initialize fuzzy sets for age
	ageYoung := FuzzySet{Name: "Young", Func: fuzzy.NewTriangle(0, 20, 40)}
	ageMiddle := FuzzySet{Name: "Middle", Func: fuzzy.NewTriangle(20, 40, 60)}
	ageOld := FuzzySet{Name: "Old", Func: fuzzy.NewTriangle(40, 60, 100)}

	// Initialize fuzzy sets for income
	incomeLow := FuzzySet{Name: "Low", Func: fuzzy.NewTriangle(0, 20000, 40000)}
	incomeMedium := FuzzySet{Name: "Medium", Func: fuzzy.NewTriangle(20000, 40000, 60000)}
	incomeHigh := FuzzySet{Name: "High", Func: fuzzy.NewTriangle(40000, 60000, 100000)}

	// Create a slice of all fuzzy sets
	allFuzzySets := []FuzzySet{ageYoung, ageMiddle, ageOld, incomeLow, incomeMedium, incomeHigh}

	// Define a fuzzy rule set
	ruleSet := []fuzzy.Rule{
		{Input: []fuzzy.Value{ageYoung.Func.Eval(30), incomeLow.Func.Eval(30000)}, Output: []fuzzy.Value{0.5}},
		{Input: []fuzzy.Value{ageMiddle.Func.Eval(45), incomeMedium.Func.Eval(45000)}, Output: []fuzzy.Value{0.7}},
		{Input: []fuzzy.Value{ageOld.Func.Eval(65), incomeHigh.Func.Eval(70000)}, Output: []fuzzy.Value{0.9}},
		// Add more rules as needed
	}

	// Create a fuzzy system
	fuzzySystem := fuzzy.NewSystem(ruleSet, allFuzzySets)

	// Example patient data
	patientAge := 50
	patientIncome := 55000

	// Calculate the fuzzy membership values for each input
	fuzzyAge := ageYoung.Func.Eval(patientAge)
	fuzzyIncome := incomeLow.Func.Eval(patientIncome)

	// Perform fuzzy inference
	output, _ := fuzzySystem.Infer([]fuzzy.Value{fuzzyAge, fuzzyIncome})

	// Determine the highest membership value for the output
	maxMembership := 0.0
	var segmentName string
	for _, outputSet := range output {
		if outputSet.Value > maxMembership {
			maxMembership = outputSet.Value
			segmentName = outputSet.Set.Name
		}
	}

	// Display the segment name
	fmt.Println("Patient Segment:", segmentName)
}
