package main

import (
	"fmt"
	"testing"
)

func BenchmarkStringInterpolation(b *testing.B) {
	var user = User{
		Name:     "Alice",
		Age:      30,
		Rating:   4.5,
		Favorite: "Blue",
	}

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf(
			"User: %s\nAge: %d\nRating: %.1f\nFavorite Color: %s",
			user.Name,
			user.Age,
			user.Rating,
			user.Favorite,
		)
	}
}

func BenchmarkFmtSprintf(b *testing.B) {
	var user = User{
		Name:     "Alice",
		Age:      30,
		Rating:   4.5,
		Favorite: "Blue",
	}

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf(
			"User: %s\nAge: %d\nRating: %.1f\nFavorite Color: %s",
			user.Name,
			user.Age,
			user.Rating,
			user.Favorite,
		)
	}
}
