package main

import (
	"fmt"
)

type User struct {
	Name     string
	Age      int
	Rating   float32
	Favorite string
}

func main() {
	user := User{
		Name:     "Alice",
		Age:      30,
		Rating:   4.5,
		Favorite: "Blue",
	}

	str := fmt.Sprintf(
		"User: %s\nAge: %d\nRating: %.1f\nFavorite Color: %s",
		user.Name,
		user.Age,
		user.Rating,
		user.Favorite,
	)
	fmt.Println(str)
}
