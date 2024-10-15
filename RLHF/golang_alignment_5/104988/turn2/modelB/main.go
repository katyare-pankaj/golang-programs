package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

// Example function with multiple paths
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}

	// Add a check for negative divisor
	if b < 0 {
		return 0, fmt.Errorf("division by a negative number is not allowed")
	}

	return a / b, nil
}

const (
	IntMax = int(^uint(0) >> 1)
	IntMin = -IntMax - 1
)

type User struct {
	Name string `valid:"required,length(3|20)"`
	Age  uint   `valid:"required,min(18)"`
}

func main() {
	user := User{Name: "John", Age: 17}
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println("Validation error:", err)
	}
}
