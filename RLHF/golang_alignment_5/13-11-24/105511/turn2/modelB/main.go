package main

import (
	"fmt"
	"math/rand"
)

// Homomorphic encryption scheme parameters
const PLAINTEXT_MODULO = 1000
const ENCRYPTION_NOISE_MAX = 100

func generateRandomNoise() int {
	return rand.Intn(ENCRYPTION_NOISE_MAX) + 1
}
func encrypt(plaintext int) int {
	noise := generateRandomNoise()
	return plaintext + noise
}
func decrypt(ciphertext int) int {
	return (ciphertext - generateRandomNoise()) % PLAINTEXT_MODULO
}
func homomorphicAdd(ciphertext1 int, ciphertext2 int) int {
	return ciphertext1 + ciphertext2
}
func main() {
	rand.Seed(42)
	plaintext1 := 20
	plaintext2 := 30
	fmt.Println("Plaintexts: ", plaintext1, plaintext2)
	ciphertext1 := encrypt(plaintext1)
	ciphertext2 := encrypt(plaintext2)
	fmt.Println("Ciphertexts: ", ciphertext1, ciphertext2)
	homomorphicSum := homomorphicAdd(ciphertext1, ciphertext2)
	decryptedSum := decrypt(homomorphicSum)
	expectedSum := plaintext1 + plaintext2
	fmt.Println("Homomorphic Sum: ", homomorphicSum)
	fmt.Println("Decrypted Sum: ", decryptedSum)
	fmt.Println("Expected Sum: ", expectedSum)
	if decryptedSum == expectedSum {
		fmt.Println("Decryption successful!")
	} else {
		fmt.Println("Decryption failed!")
	}
}
