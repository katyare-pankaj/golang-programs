package main

import (
	"crypto/subtle"
	"fmt"
)

func main() {
	secret := []byte("superSecretPassword")
	password := []byte("guessMe")

	matches := subtle.ConstantTimeCompare(secret, password)
	fmt.Println("Match:", matches == 1) // Output: Match: false
}
