// Refactored code
package main

import (
	"fmt"
)

func calculateScore(user, competency string, score int) {
	fmt.Println(fmt.Sprintf("User %s has competency %s with score %d", user, competency, score))
}
func main() {
	calculateScore("Alice", "Leadership", 8)
	calculateScore("", "Management", 7)
}
