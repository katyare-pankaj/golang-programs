package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Generate a self-signed certificate for the server
	serverCert, serverKey, err := generateCertificate("server.local")
	if err != nil {
		log.Fatal(err)
	}

	// Create a TLS config with the server certificate and key
	serverTLSConfig := &tls.Config{
		Certificates: []tls.Certificate{{
			Certificate: serverCert,
			PrivateKey:  serverKey,
		}},
		ClientAuth: tls.RequestClientCert,
	}

	// Request client certificates and verify them
	serverTLSConfig.ClientCAs = x509.NewCertPool()
	pemCABytes := []byte(`-----BEGIN CERTIFICATE-----
MIICmzCCAh8CAQAwDQYJKoZIhvcNAQELBQAwFjEUMBIGA1UEAxMLc2VydmVyLmxv
Y2FsMB4XDTIzMDcxNDE2MjM1OFoXDTI0MDcxNDE2MjM1OFowFjEUMBIGA1UEAxML
c2VydmVyLmxvY2FsMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtH/s
2y795z7MhgZm9o7FbXqK/JHv2G2EbePd0Gc5pXq36z/Rrl24dF0o3y3G30Ht4Jdq
VK8930Vd6F0/3Pqxz24z34LzXtN57+9B0z+YZBwf9206R0x50SQ7i8BJnK3Z63Bf
Hm0y4TbE1N1Q5Np1PQ7hG4JW5sNxGj/ZU9F8yzw63BXF9+4QlY+iBZ3Lnw1v+Kqw
vAqYsT1Xhj+d03SX9yp4hGgxWW0e23qk7ZiGjA7z0rBqjTjOv5GzO67OqbNwkW+L
wM+Uow9i+22Dny5w4e4zpS85x11lRqd7H1XvZbJh/H6xE+68EwIDAQABo4IBqzCC
AacwDgYDVR0PAQH/BAQDAgeAMBMGA1UdJQQMMAoGA1UdAgQGMA8wHQYDVR0OBBYE
FITsF2x0bUa7XtB1vhOgvG3Hr7W0MAkGA1UdEwQCMAAwHwYDVR0jBBgwFoAUhOwX
bHVtRrtelHW+E6C8bcets7QwEwYDVR0lBAwwCgYEVQUDBAYwFjEUMBIGA1UEAxML
c2VydmVyLmxvY2FsMB8GA1UdIwQYMBYGFBRELBdsdG1Gu17Qdb4ToLxtx6+1tDAd
BgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwGQYDVR0RBBQwEoIMc2VydmVy
LmxvY2FsMIIBDAKBgQC/b/xMv/XzjzB6CcJXZ+g/5694L5EJxjV83L1v+yWr7/UW
7X4v5cXH8b63n0Z0BxVg24kZv/fQ2H9ZcPk3iXjV8602P4vjECcJrRbEGzYUzUi5
Y0HpN9w5i8R20rI+X/Xy01V8v3x4Hx3NzTR3Bp4QP7Gh81R2T6Z8Z/T4k9Xe6C5H
tAIBADAOBgNVHQ8BAf8EBAMCAYYwDwYDVR0TAQH/BAUDAgEGMB0GA1UdDgQWBBSE
7BdsdG1Gu17Qdb4ToLxtx6+1tDAPBgNVHRMBAf8EBTADAQH/MAkGA1UdEwQCMAAw
-----END CERTIFICATE-----`)
	if ok := serverTLSConfig.ClientCAs.AppendCertsFromPEM(pemCABytes); !ok {
		log.Fatal("Failed to add client CA certificate")
	}

	// Serve the HTTP handler with TLS
	server := &http.Server{
		Addr:      ":8443",
		Handler:   http.HandlerFunc(helloHandler),
		TLSConfig: serverTLSConfig,
	}

	log.Println("Server listening on 8443")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, client! You are authenticated: %v\n", r.TLS.PeerCertificates)
}

func generateCertificate(commonName string) ([][]byte, []byte, error) {
	// Generate a new RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Create a new certificate template
	template := x509.Certificate{
		SerialNumber: 1,
		Subject:      pkix.Name{CommonName: commonName},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(1 * time.Year),
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraints: x509.BasicConstraints{
			CA: false,
		},
	}

	// Sign the certificate with the private key
	derCertificate, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, err
	}

	// Return the certificate in PEM format
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derCertificate})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	return []byte{certPEM}, keyPEM, nil
}
