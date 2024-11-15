package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Service2 struct
type Service2 struct {
	publicKey *rsa.PublicKey
}

func main() {
	// Load Service1's public key for verification (In a production environment, use a trusted certificate)
	publicKeyPEM, err := ioutil.ReadFile("service1.pub")
	if err != nil {
		log.Fatal(err)
	}

	publicKey, err := parsePublicKey(publicKeyPEM)
	if err != nil {
		log.Fatal(err)
	}

	s := &Service2{publicKey: publicKey}

	r := mux.NewRouter()
	r.HandleFunc("/receive-message", s.receiveMessage).Methods("POST")
	log.Println("Service2 listening on 8082")
	http.ListenAndServe(":8082", r)
}

func (s *Service2) receiveMessage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encryptedMessage := []byte(r.FormValue("encryptedMessage"))
	signature := []byte(r.FormValue("signature"))

	// Decrypt the message (Here, we'll use the simple XOR decryption)
	decryptedMessage := decrypt(encryptedMessage)

	// Verify the signature using Service1's public key
	if !verifySignature(s.publicKey, decryptedMessage, signature) {
		http.Error(w, "Signature verification failed", http.StatusUnauthorized)
		return
	}

	fmt.Println("Received and decrypted message from Service1:")
	fmt.Println(string(decryptedMessage))
}

func decrypt(encryptedMessage []byte) []byte {
	// Simple XOR decryption for demonstration
	decrypted := make([]byte, len(encryptedMessage))
	for i := range encryptedMessage {
		decrypted[i] = encryptedMessage[i] ^ 0x42 // Use the same key as encryption
	}
	return decrypted
}

func verifySignature(publicKey *rsa.PublicKey, message []byte, signature []byte) bool {
	h := sha256.New()
	h.Write(message)
	hash := h.Sum(nil)

	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash, signature) == nil
}

func parsePublicKey(keyPEM []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(keyPEM)
	if block == nil {
		return nil, fmt.Errorf("invalid PEM block")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	// Type assertion to *rsa.PublicKey
	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return rsaPubKey, nil
}
