package main

import (
	"fmt"
	"math/big"

	"github.com/twopir/paillier"
)

func main() {
	// Generate key pair
	priv, pub := paillier.GenerateKeyPair(2048)

	// Original message
	msg := new(big.Int).SetString("42", 10)

	// Encrypt message
	ciphertext, _ := paillier.Encrypt(pub, msg)

	// Perform some operations on the ciphertext
	// In real case, the server would perform operations on the game state
	product, _ := paillier.Multiply(pub, ciphertext, ciphertext)

	// Decrypt the result
	result, _ := paillier.Decrypt(priv, product)

	fmt.Println("Original message:", msg)
	fmt.Println("Encrypted ciphertext:", ciphertext)
	fmt.Println("Decrypted result:", result)
}
