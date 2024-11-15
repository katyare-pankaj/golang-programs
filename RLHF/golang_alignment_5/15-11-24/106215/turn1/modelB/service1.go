package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Generate a new RSA key pair for Service1 (In a production environment, use a key management system)
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	s := &Service1{privateKey: privateKey}

	r := mux.NewRouter()
	r.HandleFunc("/send-message", s.sendMessage).Methods("POST")
	log.Println("Service1 listening on 8081")
	http.ListenAndServe(":8081", r)
}

func (s *Service1) sendMessage(w http.ResponseWriter, r *http.Request) {
	// Sample message to be sent to Service2
	message := []byte("Hello, Service2! This is a secure message.")

	// Encrypt the message (Here, we'll use simple XOR for demonstration purposes. In practice, use a proper encryption scheme.)
	encryptedMessage := encrypt(message)

	// Generate digital signature using Service1's private key
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, encryptedMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"encryptedMessage": encryptedMessage,
		"signature":        signature,
	}

	fmt.Println("Sending encrypted message to Service2:")
	fmt.Println(response)
}

func encrypt(message []byte) []byte {
	// Simple XOR encryption for demonstration
	encrypted := make([]byte, len(message))
	for i := range message {
		encrypted[i] = message[i] ^ 0x42 // Use a random key in a real implementation
	}
	return encrypted
}
