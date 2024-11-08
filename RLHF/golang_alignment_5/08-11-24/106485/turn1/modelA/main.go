package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"

	"github.com/google/uuid"
)

// Property represents a property in the system
type Property struct {
	ID       uuid.UUID `json:"id"`
	Address  string    `json:"address"`
	NumUnits int       `json:"num_units"`
	// Add more fields as needed
	privateKey *rsa.PrivateKey `json:"-"` // Private key for data encryption
}

// NewProperty creates a new Property with a generated ID and private key
func NewProperty(address string, numUnits int) *Property {
	id, _ := uuid.NewRandom()
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return &Property{
		ID:         id,
		Address:    address,
		NumUnits:   numUnits,
		privateKey: privateKey,
	}
}

// EncryptData encrypts sensitive data using the property's private key
func (p *Property) EncryptData(data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, &p.privateKey.PublicKey, data)
}

// DecryptData decrypts data using the property's private key
func (p *Property) DecryptData(encryptedData []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, p.privateKey, encryptedData)
}

type PropertyManager struct {
	properties map[uuid.UUID]*Property
}

func NewPropertyManager() *PropertyManager {
	return &PropertyManager{
		properties: make(map[uuid.UUID]*Property),
	}
}

func (pm *PropertyManager) AddProperty(p *Property) {
	pm.properties[p.ID] = p
}

func (pm *PropertyManager) GetPropertyByID(id uuid.UUID) *Property {
	return pm.properties[id]
}

func main() {
	// Create a property manager
	pm := NewPropertyManager()

	// Create a new property
	newProp := NewProperty("123 Main St, Anytown, USA", 5)

	// Encrypt sensitive data (e.g., lease agreements)
	sensitiveData := []byte("Confidential lease details go here")
	encryptedData, err := newProp.EncryptData(sensitiveData)
	if err != nil {
		log.Fatalf("Error encrypting data: %v", err)
	}

	// Store the encrypted data in the property manager
	pm.AddProperty(newProp)

	// Later, retrieve the property and decrypt data as needed
	propertyID := newProp.ID
	retrievedProp := pm.GetPropertyByID(propertyID)
	if retrievedProp == nil {
		log.Fatal("Property not found.")
	}

	decryptedData, err := retrievedProp.DecryptData(encryptedData)
	if err != nil {
		log.Fatalf("Error decrypting data: %v", err)
	}

	fmt.Println("Decrypted Data:", string(decryptedData))
}
