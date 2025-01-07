package main

import (
	"errors"
	"fmt"
)

// Sentinel value for string
const invalidString = ""

// Sentinel value for int
const invalidValue = -1

func GetUsername() (string, error) {
	return invalidString, errors.New("Cannot retrieve username")
}

func GetAge() (int, error) {
	return invalidValue, errors.New("Age not available")
}

func main() {
	username, err := GetUsername()
	if username == invalidString || err != nil {
		fmt.Println("Username error:", err)
		return
	}
	fmt.Println("Username:", username)

	age, err := GetAge()
	if age == invalidValue || err != nil {
		fmt.Println("Age error:", err)
		return
	}
	fmt.Println("Age:", age)
}
