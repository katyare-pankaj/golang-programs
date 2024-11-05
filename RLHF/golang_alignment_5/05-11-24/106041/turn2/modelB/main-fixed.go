package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string   `json:"name"`
	Address *Address `json:"address"`
}

type Address struct {
	Street string  `json:"street"`
	Person *Person `json:"-"`
}

// Custom Marshaler for Address to avoid circular reference
func (a *Address) MarshalJSON() ([]byte, error) {
	type alias Address // Define an alias for Address to avoid recursion
	// Create a copy of the struct without the circular reference
	al := alias{
		Street: a.Street,
	}

	// Include PersonID derived from the Person's Name
	return json.Marshal(struct {
		Street   string `json:"street"`
		PersonID string `json:"person_id"`
	}{
		Street:   al.Street,
		PersonID: fmt.Sprintf("person_%s", a.Person.Name),
	})
}

func newPerson(name, street string) (*Person, *Address) {
	p := &Person{Name: name}
	a := &Address{Street: street, Person: p}
	p.Address = a
	return p, a
}

func main() {
	person, address := newPerson("Bob", "Elm Street 456")

	data, err := json.Marshal(person)
	if err != nil {
		log.Fatal("Error marshaling data:", err)
	}
	fmt.Println(string(data))

	data1, err := json.Marshal(address)
	if err != nil {
		log.Fatal("Error marshaling data:", err)
	}
	fmt.Println(string(data1))
}
