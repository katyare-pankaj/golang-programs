// riskassessment.go

package main

import (
	"fmt"
)

func assessRisk(severity int, likelihood int) string {
	if severity >= 7 || likelihood >= 7 {
		return "High"
	} else if severity >= 4 || likelihood >= 4 {
		return "Medium"
	} else {
		return "Low"
	}
}

func main() {
	fmt.Println(assessRisk(5, 6)) // Output: Medium
}
