package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Simulate two nodes in the distributed system
	node1 := "Node1"
	node2 := "Node2"

	// Message to be transmitted between nodes
	message := "Secret information"

	// Generate a random key for encryption
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	// Encrypt the message using AES-GCM
	ciphertext := encrypt(message, key)

	// Transmit the encrypted message from Node1 to Node2
	transmittedCiphertext := ciphertext

	// Receive the encrypted message at Node2 and decrypt it
	decryptedMessage := decrypt(transmittedCiphertext, key)

	fmt.Println("Node1:", message)
	fmt.Println("Node2:", decryptedMessage)
}

func encrypt(message string, key []byte) []byte {
	// Encrypt the message using AES-GCM
	plaintext := []byte(message)
	aesgcm, err := newAESGCM(key)
	if err != nil {
		fmt.Println("Error creating AES-GCM cipher:", err)
		return nil
	}
	nonce := make([]byte, 12)
	_, err = rand.Read(nonce)
	if err != nil {
		fmt.Println("Error generating nonce:", err)
		return nil
	}
	ciphertext, err := aesgcm.Seal(nil, nonce, plaintext, nil)
	if err != nil {
		fmt.Println("Error encrypting message:", err)
		return nil
	}

	// Add randomness to the ciphertext before transmission
	randomData := make([]byte, 16)
	_, err = rand.Read(randomData)
	if err != nil {
		fmt.Println("Error generating random data:", err)
		return nil
	}
	obfuscatedCiphertext := append(randomData, ciphertext...)
	return obfuscatedCiphertext
}

func decrypt(ciphertext []byte, key []byte) string {
	// Remove randomness from the ciphertext before decryption
	randomData := ciphertext[:16]
	ciphertext = ciphertext[16:]

	// Decrypt the message using AES-GCM
	aesgcm, err := newAESGCM(key)
	if err != nil {
		fmt.Println("Error creating AES-GCM cipher:", err)
		return ""
	}
	nonce := make([]byte, 12)
	_, err = rand.Read(nonce)
	if err != nil {
		fmt.Println("Error generating nonce:", err)
		return ""
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println("Error decrypting message:", err)
		return ""
	}

	return string(plaintext)
}

func newAESGCM(key []byte) (cipher.AEAD, error) {
	hash := sha256.New()
	_, err := hash.Write(key)
	if err != nil {
		return nil, err
	}
	key = hash.Sum(nil)[:32]
	return cipher.NewGCM(aes.NewCipher(key))
}
