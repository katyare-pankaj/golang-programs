package main

import (
	"crypto/rand"
	"fmt"
	"strings"
)

func transpose(text string, key int) string {
	cipherText := ""
	for i := 0; i < len(text); i += key {
		for j := 0; j < key; j++ {
			idx := i + j
			if idx < len(text) {
				cipherText += string(text[idx])
			}
		}
	}
	return cipherText
}

func deTranspose(cipherText string, key int) string {
	text := ""
	numColumns := len(cipherText) / key
	for j := 0; j < numColumns; j++ {
		for i := 0; i < key; i++ {
			idx := j + i*numColumns
			if idx < len(cipherText) {
				text += string(cipherText[idx])
			}
		}
	}
	return text
}

func main() {
	// Original plain text
	plainText := "Hello, this is a secret message!"
	fmt.Println("Original Text: ", plainText)

	// Secure Transposition with Random Key
	keyLength := 8
	keyBytes := make([]byte, keyLength)
	_, err := rand.Read(keyBytes)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}
	key := string(keyBytes)
	fmt.Println("Secure Transposition Key: ", key)

	cipherText := transpose(plainText, keyLength)
	fmt.Println("Encrypted Text: ", cipherText)

	// Decryption using deTranspose function
	decryptedText := deTranspose(cipherText, keyLength)
	fmt.Println("Decrypted Text: ", decryptedText)

	// Check if decryption is successful
	if strings.EqualFold(plainText, decryptedText) {
		fmt.Println("Decryption successful!")
	} else {
		fmt.Println("Decryption failed!")
	}
}
