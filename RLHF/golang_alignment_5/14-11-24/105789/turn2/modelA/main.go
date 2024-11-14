package main

import (
	"crypto/rand"
	"fmt"
)

const (
	modulus    = 1009 // The modulus N, should be a prime number
	noiseLevel = 100  // The maximum noise level allowed
)

// Encrypts an integer using random noise
func encrypt(x int) int {
	noise := rand.Intn(noiseLevel)
	return (x + noise) % modulus
}

// Decrypts an encrypted integer
func decrypt(c int) int {
	for i := 0; i < modulus; i++ {
		if (i+noiseLevel)%modulus == c {
			return i
		}
	}
	return -1 // Should never be reached
}

// Performs homomorphic addition on two encrypted integers
func homomorphicAdd(c1, c2 int) int {
	return (c1 + c2) % modulus
}

func main() {
	// Original integers to add
	x1 := 42
	x2 := 58

	// Encrypt the original integers
	c1 := encrypt(x1)
	c2 := encrypt(x2)

	// Perform homomorphic addition on the encrypted integers
	c3 := homomorphicAdd(c1, c2)

	// Decrypt the result
	result := decrypt(c3)

	// Output results
	fmt.Println("Original value x1:", x1)
	fmt.Println("Original value x2:", x2)
	fmt.Println("Encrypted value c1:", c1)
	fmt.Println("Encrypted value c2:", c2)
	fmt.Println("Encrypted addition result c3:", c3)
	fmt.Println("Decrypted result:", result)
}
