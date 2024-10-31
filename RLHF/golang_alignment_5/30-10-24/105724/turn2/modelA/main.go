package modelA

import (
	"fmt"
)

// Step 1: Encapsulate user data
type User struct {
	name  string
	email string
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func main() {
	user := &User{name: "Alice"}
	fmt.Println("User Name:", user.GetName())

	// Modifying user name encapsulated change
	user.SetName("Bob")
	fmt.Println("Updated User Name:", user.GetName())
}
