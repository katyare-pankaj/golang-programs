package main

import "fmt"

// User struct encapsulates user data
type User struct {
	name    string
	email   string
	address Address
}

// Address struct encapsulates address data
type Address struct {
	street  string
	city    string
	country string
}

func (u *User) SetAddress(street, city, country string) {
	u.address = Address{street, city, country}
}

func main() {
	user := &User{name: "John Doe", email: "johndoe@example.com"}

	// Set user's address using encapsulated method
	user.SetAddress("123 Main St", "New York", "USA")

	fmt.Println("User Name:", user.name)
	fmt.Println("User Address:", user.address)
}
