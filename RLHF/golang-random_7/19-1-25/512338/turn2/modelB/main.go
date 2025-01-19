package main  
import (  
    "crypto/aes"  
    "crypto/cipher"  
    "crypto/rand"  
    "encoding/base64"  
    "errors"  
    "fmt"  
    "io"  
    "log"  
    "net/http"  
    "strings"  
    "sync"  
)  
  
type PrivacyLevel int  
const (  
    Public PrivacyLevel = iota  
    Private  
    Sensitive  
)  
  
var privacyLevels = map[string]PrivacyLevel{  
    "user_id":      Sensitive,  
    "email":        Private,  
    "query":       Public,  
    "custom_param": Sensitive,  
}  
  
var encryptionKey = "your-32-byte-encryption-key-here" // Replace this with a secure key

var mu sync.Mutex // To protect concurrent access to the encryption/decryption

func init() {
    if len(encryptionKey) != 32 {
        log.Fatal("Encryption key must be 32 bytes long")
    }
}

// Encrypt data for privacy
func encrypt(data string) (string, error) {
    mu.Lock()
    defer mu.Unlock()
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
    return base64.URLEncoding.EncodeToString(encryptedData), nil
}
  
// Decrypt data 
func decrypt(encryptedData string) (string, error) {  
    mu.Lock()
    defer mu.Unlock()
    data, err := base64.URLEncoding.DecodeString(encryptedData)  
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
    if len(data) < nonceSize {  
        return "", errors.New("invalid data size")  
    }  
  
    nonce, ciphertext := data[:nonceSize], data[nonceSize:]  
    decryptedData, err := gcm.Open(nil, nonce, ciphertext, nil)  
    if err != nil {  
        return "", err  
    }  
  
    return string(decryptedData), nil  
}  
  
func getPrivacyLevel(param string) PrivacyLevel {  
    if privacyLevel, ok := privacyLevels[strings.ToLower(param)]; ok {  
        return privacyLevel  
    }  
    return Public // Default privacy level for unknown parameters  
}  
  
func isSensitive(privacyLevel PrivacyLevel) bool {  
    return privacyLevel >= Sensitive  
}  
  
func handleQueryParams(w http.ResponseWriter, r *http.Request) {  
    params := r.URL.Query()  
    for key, value := range params {  
        privacyLevel := getPrivacyLevel(key)  
        encryptedValue, err := encrypt(value[0])  
        if err != nil {  
            log.Printf("Error encrypting value for %s: %v", key, err)  
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)  
            return  
        }