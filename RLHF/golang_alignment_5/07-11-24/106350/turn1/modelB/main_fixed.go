package main

import (
	"fmt"
	"strings"
)

// Transpose function for encryption
func transpose(text string, key int) string {

	numRows := (len(text) + key - 1) / key

	cipherText := ""
	// Read the text column by column
	for j := 0; j < key; j++ {
		for i := 0; i < numRows; i++ {
			idx := i*key + j
			if idx < len(text) {
				cipherText += string(text[idx])
			}
		}
	}
	return cipherText
}

// DeTranspose function for decryption
func deTranspose(cipherText string, key int) string {
	numRows := len(cipherText) / key
	if len(cipherText)%key != 0 {
		numRows++
	}

	// Create an empty matrix to hold the characters
	matrix := make([][]rune, numRows)
	for i := range matrix {
		matrix[i] = make([]rune, key)
	}

	// Fill the matrix column by column from the cipherText
	idx := 0
	for j := 0; j < key; j++ {
		for i := 0; i < numRows; i++ {
			if idx < len(cipherText) {
				matrix[i][j] = rune(cipherText[idx])
				idx++
			}
		}
	}

	// Read the matrix row by row to reconstruct the original text
	text := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < key; j++ {
			if matrix[i][j] != 0 {
				text += string(matrix[i][j])
			}
		}
	}
	return text
}

func main() {
	// Original plain text
	plainText := "Hello, this is a secret message!"
	fmt.Println("Original Text: ", plainText)

	// Generate a random key length (between 2 and 10 for example)
	keyLength := 8

	// Secure Transposition with Random Key Length (key length is an integer, not a byte slice)
	fmt.Println("Secure Transposition Key Length: ", keyLength)

	// Encrypt the text using the transpose function
	cipherText := transpose(plainText, keyLength)
	fmt.Println("Encrypted Text: ", cipherText)

	// Decrypt the text using deTranspose function
	decryptedText := deTranspose(cipherText, keyLength)
	fmt.Println("Decrypted Text: ", decryptedText)

	// Check if decryption is successful
	if strings.EqualFold(plainText, decryptedText) {
		fmt.Println("Decryption successful!")
	} else {
		fmt.Println("Decryption failed!")
	}
}
