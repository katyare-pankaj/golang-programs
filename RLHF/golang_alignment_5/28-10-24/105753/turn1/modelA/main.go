package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

func encrypt(key []byte, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// IV needs to be unique for each encryption
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], text)

	return ciphertext, nil
}

func decrypt(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

func main() {
	// Example key. In a real application, you should manage this securely
	key := []byte("32-byte-long-secret-key-goes-here-change-this-please")

	dataToEncrypt := []byte("Sensitive information goes here")

	// Encrypt
	ciphertext, err := encrypt(key, dataToEncrypt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encrypted data: %x\n", ciphertext)

	// Decrypt
	plaintext, err := decrypt(key, ciphertext)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decrypted data: %s\n", plaintext)
}
