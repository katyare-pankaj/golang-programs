package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/hashicorp/vault/api"
)

func encrypt(key []byte, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], text)

	return ciphertext, nil
}

func decrypt(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

func getKeyFromVault(vaultAddr, vaultToken, keyPath string) ([]byte, error) {
	client, err := api.NewClient(&api.Config{
		Address: vaultAddr,
	})
	if err != nil {
		return nil, err
	}

	client.SetToken(vaultToken)

	secret, err := client.Logical().Read(keyPath)
	if err != nil {
		return nil, err
	}

	if secret == nil {
		return nil, fmt.Errorf("key not found at path: %s", keyPath)
	}

	keyData, ok := secret.Data["key"]
	if !ok {
		return nil, fmt.Errorf("key data not found in secret")
	}

	return []byte(keyData.(string)), nil
}

func main() {
	// Vault configuration
	vaultAddr := "http://localhost:8200" // Change this to your Vault address
	vaultToken := "s.exampleVaultToken"   // Change this to your Vault token
	keyPath := "sportsapp/data-encryption-key" // Change this to the path where your key is stored in Vault

	// Data to encrypt
	userData := []byte("This is sensitive user data")

	// Retrieve the encryption key from Vault
	key, err := getKeyFromVault(vaultAddr, vaultToken, keyPath)
	if err != nil {
		panic(err)
	}

	// Encrypt the data
	ciphertext, err := encrypt(key, userData)
	if err != nil {
		panic(err)
	}

	// Base64 encode the ciphertext for storage or transmission
	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Printf("Encrypted data: %s\n", encodedCiphertext)

	// ...

	// To decrypt, retrieve the key from Vault again (in a real app, you should cache the key)
	// key, err = getKeyFromVault(vaultAddr, vaultToken, keyPath)
	// if err != nil {
	// 	panic(err)
	// }

	// Decrypt the data
	// ciphertext, _ = base64.StdEncoding.DecodeString(encodedCiphertext)
	// plaintext, err := decrypt(key, ciphertext)