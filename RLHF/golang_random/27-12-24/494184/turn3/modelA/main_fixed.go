package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	// Define custom error types for better readability
	ErrInputInvalid    = errors.New("input is invalid")
	ErrDatabaseFailure = errors.New("database operation failed")
)

func step1(data string) error {
	if len(data) == 0 {
		return fmt.Errorf("%w: data is empty", ErrInputInvalid)
	}
	return nil
}

func step2(data string) error {
	// Simulate a database operation
	if rand.Intn(2) == 0 {
		return fmt.Errorf("%w: could not insert data into database", ErrDatabaseFailure)
	}
	return nil
}

func step3(data string) error {
	// Simulate an additional processing step
	if len(data) < 3 {
		return fmt.Errorf("%w: data length must be at least 3", ErrInputInvalid)
	}
	return nil
}

func process(data string) error {
	err := step1(data)
	if err != nil {
		return err
	}

	err = step2(data)
	if err != nil {
		return err
	}

	return step3(data)
}

func main() {
	data := "hello"
	err := process(data)
	if err != nil {
		fmt.Printf("Process failed: %v\n", err)
		if isInputInvalid(err) {
			fmt.Println("Input is invalid.")
		} else if errors.Is(err, ErrDatabaseFailure) {
			fmt.Println("Database operation failed.")
		}
	} else {
		fmt.Println("Process completed successfully!")
	}
}

func isInputInvalid(err error) bool {
	return errors.Is(err, ErrInputInvalid)
}
