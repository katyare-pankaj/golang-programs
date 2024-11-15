package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"tls"
)

// Simple backend microservice
func backendHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the backend microservice!")
}

func main() {
	// Generate a new RSA key pair for the server (For demonstration purposes, in a real scenario, use a trusted CA)
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Error generating RSA key: %v", err)
	}

	pubKey := &priv.PublicKey

	// Create the server's certificate template
	template := x509.Certificate{
		Subject:      pkix.Name{CommonName: "telecom-backend"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Year),
		SerialNumber: big.NewInt(1),
		DNSNames:     []string{"telecom-backend.local"},
	}

	// Self-sign the certificate for the server
	cert, err := x509.CreateCertificate(rand.Reader, &template, &template, pubKey, priv)
	if err != nil {
		log.Fatalf("Error creating self-signed certificate: %v", err)
	}

	caCert, caKey, err := generateCACert() // Create a simple CA for demonstration
	if err != nil {
		log.Fatalf("Error creating CA certificate: %v", err)
	}

	caFilePath := "ca.crt"
	certFilePath := "server.crt"
	keyFilePath := "server.key"

	saveCertificate(cert, certFilePath)
	saveCertificate(caCert, caFilePath)
	savePrivateKey(priv, keyFilePath)
	savePrivateKey(caKey, "ca.key") // save CA private key

	serverTLSConfig := &tls.Config{
		Certificates: []tls.Certificate{
			{
				Certificate: []byte(cert),
				PrivateKey:  []byte(loadPrivateKey(keyFilePath)),
			},
		},
		ClientCAs:                createClientCACertPool(),
		ClientAuth:               tls.RequireAndVerifyClientCert,
		TraceLog:                 log.New(os.Stdout, "TLS: ", log.Lshortfile),
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
	}

	// Enable TLS for the backend server
	server := &http.Server{
		Addr:      ":8443",
		Handler:   http.HandlerFunc(backendHandler),
		TLSConfig: serverTLSConfig,
	}

	log.Println("Backend microservice listening on 8443 with TLS")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func generateCACert() ([]byte, *rsa.PrivateKey, error) {
	// Generate CA key pair
	caKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	caPubKey := &caKey.PublicKey

	template := x509.Certificate{
		Subject:      pkix.Name{CommonName: "telecom-CA"},
		IsCA:         true,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Year),
		SerialNumber: big.NewInt(1),
	}

	caCert, err := x509.CreateCertificate(rand.Reader, &template, &template, caPubKey, caKey)
	if err != nil {
		return nil, nil, err
	}

	return caCert, caKey, nil
}

func saveCertificate(cert []byte, filePath string) {
	certOut, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating certificate file %s: %v", filePath, err)
	}
	defer certOut.Close()

	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert})
	log.Println("Saved certificate to", filePath)
}

func savePrivateKey(key *rsa.PrivateKey, filePath string) {
	privateKeyOut, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating private key file %s: %v", filePath, err)
	}
	defer privateKeyOut.Close()

	encryptedKey, err := x509.EncryptPKCS1PrivateKey(key, []byte("telecom-password"))
	if err != nil {
		log.Fatalf("Error encrypting private key: %v", err)
	}

	pem.Encode(privateKeyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: encryptedKey})
	log.Println("Saved private key to", filePath)
}

func loadPrivateKey(filePath string) []byte {
	privateKeyFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error loading private key file %s: %v", filePath, err)
	}
	defer privateKeyFile.Close()

	block, _ := pem.Decode(privateKeyFile)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatalf("Error decoding private key file %s: %v", filePath, err)
	}

	decryptedKey, err := x509.DecryptPKCS1PrivateKey(block.Bytes, []byte("telecom-password"))
	if err != nil {
		log.Fatalf("Error decrypting private key: %v", err)
	}

	return x509.MarshalPKCS1PrivateKey(decryptedKey)
}

func createClientCACertPool() *x509.CertPool {
	caCertPool := x509.NewCertPool()
	caCertFile, err := os.Open("ca.crt")
	if err != nil {
		log.Fatalf("Error opening CA certificate file: %v", err)
	}
	defer caCertFile.Close()
	caCertBytes, err := ioutil.ReadAll(caCertFile)
	if err != nil {
		log.Fatalf("Error reading CA certificate file: %v", err)
	}

	if !caCertPool.AppendCertsFromPEM(caCertBytes) {
		log.Fatalf("Error adding CA certificate to pool")
	}

	return caCertPool
}
