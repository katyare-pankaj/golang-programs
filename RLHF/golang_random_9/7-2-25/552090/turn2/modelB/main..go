package main  
import (  
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

// Request struct to hold the incoming message
type Request struct {  
    Message string `json:"message"`
}

// Response struct to hold the processed message
type Response struct {  
    ProcessedMessage string `json:"processed_message"`
}

func handler(w http.ResponseWriter, r *http.Request) {  
    // Ensure the request body is closed after reading
    defer r.Body.Close()

    // Decode the incoming JSON request
    var req Request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Create a context with a timeout for resource cleanup
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Open a file for logging
    logFile, err := os.OpenFile("serverless.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)  
    if err != nil {  
        http.Error(w, "Could not open log file", http.StatusInternalServerError)
        return
    }
    // Ensure the log file is closed after use
    defer logFile.Close()

    // Set the log output to the log file
    log.SetOutput(logFile)

    // Simulate a resource that needs cleanup (e.g., a database connection)
    resource, err := acquireResource(ctx)
    if err != nil {
        log.Printf("Error acquiring resource: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    // Ensure the resource is released after processing
    defer resource.Close()

    // Log the incoming message
    log.Printf("Received message: %s", req.Message)

    // Simulate processing the message using the acquired resource
    processedMessage, err := processMessage(ctx, resource, req.Message)
    if err != nil {
        log.Printf("Error processing message: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Respond with the processed message
    response := Response{ProcessedMessage: processedMessage}  
    w.Header().Set("Content-Type", "application/json")  
    json.NewEncoder(w).Encode(response)

    // Log the response
    log.Printf("Sent response: %s", processedMessage)
}

// Simulate acquiring a resource (e.g., database connection)
func acquireResource(ctx context.Context) (resource *os.File, err error) {
    // Simulate a delay for resource acquisition (can be omitted)
    time.Sleep(1 * time.Second)

    // Create a temporary file for the resource
    resource, err = os.CreateTemp("", "resource.*")
    if err != nil {
        return nil, fmt.Errorf("error creating temporary resource: %w", err)
    }

    return resource, nil
}

// Simulate processing the message using the acquired resource
func processMessage(ctx context.Context, resource *os.File, message string) (string, error) {
    // Simulate processing the message using the resource
    _, err := resource.WriteString(fmt.Sprintf("Processed: %s\n", message))
    if err != nil {
        return "", fmt.Errorf("error writing to resource: %w", err)
    }

    // Simulate a delay for processing (can be omitted)
    time.Sleep(2 * time.Second)

    // Return the processed message