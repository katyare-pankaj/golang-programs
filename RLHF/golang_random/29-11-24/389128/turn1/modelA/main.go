package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/29-11-24/389128/turn1/modelA/feature"
	"log"
)

func main() {
	if feature.GetFlag("ENABLE_NEW_FEATURE") {
		log.Println("New feature is enabled.")
		newFeatureFunctionality()
	} else {
		log.Println("New feature is disabled.")
	}
}

func newFeatureFunctionality() {
	fmt.Println("Performing new feature operations...")
}
