package main

import (
	"fmt"
	"math/rand"
)

// Simple homomorphic encryption scheme (additive)
func encrypt(x int, N int, r int) int {

	return (x + r) % N
}

func decrypt(c int, r int, N int) int {

	return (c - r + N) % N
}

func addEncrypted(c1 int, c2 int, N int) int {

	return (c1 + c2) % N
}

func main() {
	// Set up parameters

	N := 1009
	r1 := rand.Intn(N)
	r2 := rand.Intn(N)

	// Original messages (game state data)
	x1 := 42
	x2 := 58

	// Encrypt messages
	c1 := encrypt(x1, N, r1)
	c2 := encrypt(x2, N, r2)

	// Perform homomorphic addition on encrypted values
	c3 := addEncrypted(c1, c2, N)

	// Decrypt the result
	decryptedResult := decrypt(c3, r1+r2, N)

	// Output results
	fmt.Println("Original value x1:", x1)
	fmt.Println("Original value x2:", x2)
	fmt.Println("Encrypted value c1:", c1)
	fmt.Println("Encrypted value c2:", c2)
	fmt.Println("Encrypted addition result c3:", c3)
	fmt.Println("Decrypted result:", decryptedResult)
}
