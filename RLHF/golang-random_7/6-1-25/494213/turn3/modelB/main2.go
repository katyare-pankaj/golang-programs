package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

type UserDataError struct {
	Code    int
	Message string
}

func (e *UserDataError) Error() string {
	return fmt.Sprintf("UserDataError: Code: %d, Message: %s", e.Code, e.Message)
}

func GetUserData() (user *User, age int, err error) {
	// Some complex logic that might fail
	// ...

	// If an error occurs
	return nil, 0, &UserDataError{Code: 1001, Message: "Failed to fetch user data"}
}

func main() {
	user, age, err := GetUserData()
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		if ude, ok := err.(*UserDataError); ok {
			fmt.Println("Error Code:", ude.Code)
		}
		return
	}
	fmt.Println("User:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("Additional Age:", age)
}
