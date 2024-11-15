package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"log"
	"time"
)

func main() {
	// Generate a new RSA key pair for PartyA
	partyAPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	// Generate a new RSA key pair for PartyB
	partyBPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	// Create a simple self-signed certificate for PartyA
	partyACert, err := createSelfSignedCert(partyAPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Create a simple self-signed certificate for PartyB
	partyBCert, err := createSelfSignedCert(partyBPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// PartyA encrypts a message for PartyB using PartyB's public key
	message := []byte("Hello, PartyB! This is a secure message from PartyA.")
	encryptedMessage, err := encryptMessage(partyBPublicKey(partyBCert), message)
	if err != nil {
		log.Fatal(err)
	}

	// PartyB receives the encrypted message and decrypts it using its private key
	decryptedMessage, err := decryptMessage(partyBPrivateKey, encryptedMessage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PartyA sent:", string(message))
	fmt.Println("PartyB received:", string(decryptedMessage))
}

func createSelfSignedCert(privateKey *rsa.PrivateKey) (*x509.Certificate, error) {
	template := &x509.Certificate{
		SerialNumber: 1,
		Subject:      x509.Name{CommonName: "PartyA"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Year),
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraints: &x509.BasicConstraints{
			IsCA:       false,
			MaxPathLen: nil,
		},
	}

	return x509.CreateCertificate(rand.Reader, template, template, privateKey, privateKey)
}

func partyBPublicKey(cert *x509.Certificate) *rsa.PublicKey {
	return cert.PublicKey.(*rsa.PublicKey)
}

func encryptMessage(publicKey *rsa.PublicKey, message []byte) ([]byte, error) {
	return rsa.EncryptOAEP(rand.Reader, nil, publicKey, message, []byte("example")) // Replace "example" with a secure padding string
}

func decryptMessage(privateKey *rsa.PrivateKey, encryptedMessage []byte) ([]byte, error) {
	return rsa.DecryptOAEP(rand.Reader, nil, privateKey, encryptedMessage, []byte("example")) // Replace "example" with the same secure padding string
}
