package main

import (
	"errors"
	"fmt"
)

func step1() error {
	// Simulate step 1 failure
	return errors.New("step 1 failed")
}

func step2(prevErr error) error {
	// Simulate step 2 failure
	return errors.Wrap(prevErr, "step 2 failed")
}

func step3(prevErr error) error {
	// Simulate step 3 failure
	return errors.Wrap(prevErr, "step 3 failed")
}

func multiStepProcess() error {
	err := step1()
	if err != nil {
		return err
	}

	err = step2(err)
	if err != nil {
		return err
	}

	err = step3(err)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := multiStepProcess()
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	} else {
		fmt.Println("Multi-step process completed successfully!")
	}
}
