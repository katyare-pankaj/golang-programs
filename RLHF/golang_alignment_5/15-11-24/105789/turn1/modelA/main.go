package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func main() {
	host := os.Getenv("TARGET_HOST")
	if host == "" {
		log.Fatal("TARGET_HOST environment variable not set.")
	}

	resetLightBrightness("dim", host)
	time.Sleep(1 * time.Second)
	resetLightBrightness("dim", host) // Resending the request with the same idempotency key
}

func resetLightBrightness(brightness string, host string) {
	idempotencyKey := generateIdempotencyKey()
	url := fmt.Sprintf("http://%s/api/v1/lights/brightness", host)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Add("Idempotency-Key", idempotencyKey)
	req.Header.Add("Content-Type", "application/json")

	payload := map[string]string{"brightness": brightness}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to reset brightness. Status: %s", resp.Status)
	}
	fmt.Println("Payload", payload)
	fmt.Printf("Brightness reset successfully with idempotency key: %s\n", idempotencyKey)
}

func generateIdempotencyKey() string {
	return fmt.Sprintf("%v-%v", time.Now().Unix(), uuid.New())
}
