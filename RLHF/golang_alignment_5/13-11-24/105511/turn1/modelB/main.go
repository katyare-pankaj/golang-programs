package main

import (
	"fmt"

	"github.com/zhangliangxiao/he"
)

func main() {
	// Initialize the homomorphic encryption scheme
	params := he.NewParameters(1024, 64) // Key size and plaintext modulus
	scheme := he.NewScheme(params)

	// Generate a key pair
	privateKey, publicKey := scheme.KeyGen()

	// Encrypt a game state (e.g., player score)
	plaintext := 100
	ciphertext := scheme.Encrypt(publicKey, plaintext)

	// Perform a homomorphic addition on the encrypted game state
	ciphertext2 := scheme.Add(ciphertext, ciphertext)

	// Decrypt the result
	decrypted := scheme.Decrypt(privateKey, ciphertext2)

	fmt.Println("Decrypted result:", decrypted) // Output: Decrypted result: 200
}
