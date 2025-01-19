package main  
import (  
    "context"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "io"
    "log"
    "net/http"
    "sync"
    "time"

    "github.com/gorilla/mux"
    "github.com/juju/ratelimit"
)

var (
    // Define privacy levels
    Public   PrivacyLevel = iota
    Private
    Sensitive

    // Store encryption keys for each privacy level
    encryptionKeys = map[PrivacyLevel][]byte{
        Public:   []byte("publickey123456"),
        Private:  []byte("privatekey123456"),
        Sensitive: []byte("sensitivekey123456"),
    }

    // Create a sync.Pool to manage AES cipher blocks
    cipherBlockPool = sync.Pool{
        New: func() interface{} {
            block, err := aes.NewCipher([]byte(encryptionKeys[Sensitive]))
            if err != nil {
                log.Fatalf("Error creating cipher block: %v", err)
            }
            return block
        },
    }

    // Create rate limiters for encryption and decryption operations
    encryptRateLimiter = ratelimit.NewBucketWithRate(100, 100)
    decryptRateLimiter = ratelimit.NewBucketWithRate(100, 100)
)

// PrivacyLevel is an enumeration for privacy levels
type PrivacyLevel int

// Encrypt encrypts sensitive data
func encrypt(data string) (string, error) {
    // Acquire a cipher block from the pool
    block := cipherBlockPool.Get().(cipher.Block)
    defer cipherBlockPool.Put(block)

    // Acquire a rate limit token
    if err := encryptRateLimiter.Wait(context.Background()); err != nil {
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

// Decrypt decrypts sensitive data
func decrypt(encryptedData string) (string, error) {
    // Acquire a cipher block from the pool
    block := cipherBlockPool.Get().(cipher.Block)
    defer cipherBlockPool.Put(block)
    
    // Acquire a rate limit token
    if err := decryptRateLimiter.Wait(context.Background()); err != nil {
        return "", err
    }

    data, err := base64.StdEncoding.DecodeString(encryptedData)
    if err != nil {
        return "", err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    nonceSize := gcm.NonceSize()
    if len(data) < nonceSize {
        return "", fmt.Errorf("invalid data size")
    }

    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    decryptedData, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return "", err
    }

    return string(decryptedData), nil
}

// Handle query parameters with respect to their privacy level
func queryParamHandler(w http.ResponseWriter, r *http.Request) {
    queryParams := r.URL.Query()

    // Get the privacy level from the request context
    privacyLevel, ok := r.Context().Value("privacyLevel").(PrivacyLevel)
    if !ok {
        http.Error(w, "Invalid privacy level", http.StatusBadRequest)
        return
    }

    for key, value := range queryParams {
        if len(value) == 0 {
            continue
        }

        var processedValue string

        switch privacyLevel {
        case Public:
            processedValue = value[0] // Direct usage
            fmt.Fprintf(w, "Public %s: %s\n", key, processedValue)
        case Private, Sensitive:
            // Encrypt the value
            encryptedValue, err := encrypt(value[0])
            if err != nil {
                log.Printf("Encryption error: %v", err)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
            }
            processedValue = encryptedValue
            fmt.Fprintf(w, "%s %s (Encrypted): %s\n", privacyLevel, key, processedValue)
        }
    }
}

func main() {
    r := mux.NewRouter()

    // Middleware to set the privacy level based on the route
    privacyLevelMiddleware := func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            vars := mux.Vars(r)
            privacyLevelStr := vars["privacyLevel"]

            var privacyLevel PrivacyLevel
            switch privacyLevelStr {
            case "public":
                privacyLevel = Public
            case "private":
                privacyLevel = Private
            case "sensitive":
                privacyLevel = Sensitive
            default:
                http.Error(w, "Invalid privacy level", http.StatusBadRequest)
                return
            }

            // Set the privacy level in the request context
            ctx := context.WithValue(r.Context(), "privacyLevel", privacyLevel)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }

    r.HandleFunc("/query/{privacyLevel}", privacyLevelMiddleware(queryParamHandler))

    fmt.Println("Server is listening on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}