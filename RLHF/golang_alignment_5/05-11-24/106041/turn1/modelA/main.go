package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Person struct {
	Name    string   `json:"name"`
	Address *Address `json:"address"`
}

type Address struct {
	Street   string  `json:"street"`
	Person   *Person `json:"-"` // Ignore the Person reference to break the cycle
	PersonID string  `json:"person_id"`
}

func newPerson(name, street string) (*Person, *Address) {
	p := &Person{
		Name: name,
	}
	a := &Address{
		Street:   street,
		PersonID: fmt.Sprintf("person_%s", strings.ToLower(p.Name)),
	}
	p.Address = a
	return p, a
}

func main() {
	person, address := newPerson("Alice", "Main Street 123")

	data, err := json.Marshal(person)
	if err != nil {
		log.Fatal("Error marshaling data:", err)
	}
	fmt.Println(string(data))
}
