package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)

	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	// Your data to be encrypted
	data := []byte("This is some sensitive data that needs encryption!")

	// Encryption key - Must be 16, 24, or 32 bytes long for AES
	key := []byte("32byte16byte24byte32byte32byte32byte32byte32byte")

	// Encrypt the data
	ciphertext, err := encrypt(key, data)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}

	// Display the encrypted data (base64 encoded)
	fmt.Println("Encrypted data:", base64.StdEncoding.EncodeToString(ciphertext))

	// Decrypt the data
	plaintext, err := decrypt(key, ciphertext)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	// Display the decrypted data
	fmt.Println("Decrypted data:", string(plaintext))
}
