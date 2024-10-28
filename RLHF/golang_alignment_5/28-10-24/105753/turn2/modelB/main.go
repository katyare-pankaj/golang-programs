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

	// Never use more than 2^32 random nonces with a given key for security reasons.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, text, nil)
	return ciphertext, nil
}

func decrypt(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func main() {
	// Example key. In a real application, you should manage this securely.
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
