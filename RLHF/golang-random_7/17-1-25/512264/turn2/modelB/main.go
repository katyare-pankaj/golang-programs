package main

import (
	"errors"
	"fmt"
	"os"
)

type Resource struct {
	file *os.File
}

func (r *Resource) Close() error {
	if r.file != nil {
		return r.file.Close()
	}
	return nil
}

func OpenResource() (*Resource, error) {
	f, err := os.Create("test.txt")
	if err != nil {
		return nil, err
	}
	r := &Resource{file: f}
	return r, nil
}

func DoSomething() error {
	r, err := OpenResource()
	if err != nil {
		return err
	}

	// Defer the resource close to ensure it gets executed regardless of error
	defer r.Close()

	// Simulate some work that might fail
	if true {
		return errors.New("simulated work error")
	}

	// If no error, the resource will be closed automatically when the function returns
	fmt.Println("Work completed successfully.")
	return nil
}

func main() {
	if err := DoSomething(); err != nil {
		fmt.Println("Error:", err)
	}
}
