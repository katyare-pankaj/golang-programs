package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
)

const encryptionKey = "examplekey123456" // Encryption key for AES (must be 16, 24, or 32 bytes long)

// Encrypt data for privacy
func encrypt(data string) (string, error) {
	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encryptedData := gcm.Seal(nonce, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

// Decrypt data
func decrypt(encryptedData string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("invalid data size")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	decryptedData, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedData), nil
}

// Handler to manage query parameters
func queryParamHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	for key, value := range queryParams {
		// Encrypt sensitive data
		encryptedValue, err := encrypt(value[0])
		if err != nil {
			log.Printf("Error encrypting value: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Encrypted %s: %s\n", key, encryptedValue)

		// Decrypt to show functionality
		decryptedValue, err := decrypt(encryptedValue)
		if err != nil {
			log.Printf("Error decrypting value: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Decrypted %s: %s\n", key, decryptedValue)
	}
}

func main() {
	http.HandleFunc("/query", queryParamHandler)

	fmt.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
