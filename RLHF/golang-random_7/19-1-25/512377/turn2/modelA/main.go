package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Defined privacy levels
type PrivacyLevel int

const (
	Public PrivacyLevel = iota
	Private
	Sensitive
)

const encryptionKey = "examplekey123456" // AES key for encryption (16 bytes)

var paramPrivacyLevels = map[string]PrivacyLevel{
	"username": Public,
	"email":    Private,
	"ssn":      Sensitive,
}

// Encrypt sensitive data
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

// Handle query parameters with respect to their privacy level
func queryParamHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	for key, value := range queryParams {
		if len(value) == 0 {
			continue
		}

		var processedValue string

		switch paramPrivacyLevels[key] {
		case Public:
			processedValue = value[0] // Direct usage
			fmt.Fprintf(w, "Public %s: %s\n", key, processedValue)

		case Private:
			// Hash the value using HMAC-SHA256
			h := hmac.New(sha256.New, []byte(encryptionKey))
			h.Write([]byte(value[0]))
			processedValue = base64.StdEncoding.EncodeToString(h.Sum(nil))
			fmt.Fprintf(w, "Private %s (Hashed): %s\n", key, processedValue)

		case Sensitive:
			// Encrypt the value
			encryptedValue, err := encrypt(value[0])
			if err != nil {
				log.Printf("Encryption error: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			processedValue = encryptedValue
			fmt.Fprintf(w, "Sensitive %s (Encrypted): %s\n", key, processedValue)
		}
	}
}

func main() {
	http.HandleFunc("/query", queryParamHandler)

	fmt.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
