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
	"sync"
)

// Define privacy levels
type PrivacyLevel int

const (
	Public PrivacyLevel = iota
	Private
	Sensitive
)

const encryptionKey = "examplekey123456" // Encryption key (16 bytes for AES-128)

var paramPrivacyLevels = map[string]PrivacyLevel{
	"username": Public,
	"email":    Private,
	"ssn":      Sensitive,
}

// Encrypt uses AES-GCM to encrypt sensitive data
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
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encryptedData := gcm.Seal(nonce, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

// Decrypt uses AES-GCM to decrypt sensitive data
func decrypt(encryptedData string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(encryptedData)
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
	if len(decodedData) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := decodedData[:nonceSize], decodedData[nonceSize:]
	decryptedData, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedData), nil
}

// Handler to process query parameters securely
func queryParamHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	var wg sync.WaitGroup

	results := make(chan string, len(queryParams))

	for key, value := range queryParams {
		// Capture the key and value for the goroutine
		key, value := key, value

		wg.Add(1)
		go func() {
			defer wg.Done()

			if len(value) == 0 {
				results <- fmt.Sprintf("%s: No value provided\n", key)
				return
			}

			var processedValue string
			var err error

			switch paramPrivacyLevels[key] {
			case Public:
				processedValue = value[0] // Direct usage
				results <- fmt.Sprintf("Public %s: %s\n", key, processedValue)

			case Private:
				// Assume hashing or pseudonymization is done here
				processedValue = "hashed_" + value[0] // Simplified for demonstration
				results <- fmt.Sprintf("Private %s (Hashed): %s\n", key, processedValue)

			case Sensitive:
				// Encrypt the value
				processedValue, err = encrypt(value[0])
				if err != nil {
					log.Printf("Encryption error: %v", err)
					results <- fmt.Sprintf("Sensitive %s: Error encrypting data\n", key)
					return
				}
				results <- fmt.Sprintf("Sensitive %s (Encrypted): %s\n", key, processedValue)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Fprint(w, result)
	}
}

func main() {
	http.HandleFunc("/query", queryParamHandler)

	fmt.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
