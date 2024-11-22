package main

import (
	"context"
	"log"
	"time"
)

// simulateDataFetch simulates fetching data from an external service.
// It takes a context, a timeout, and a request ID for logging.
func simulateDataFetch(ctx context.Context, requestID string, timeout time.Duration) ([]byte, error) {
	select {
	case <-time.After(timeout):
		return nil, context.DeadlineExceeded
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		// Simulate some work being done.
		time.Sleep(time.Second * 2)
		return []byte("Simulated data from external service"), nil
	}
}

func main() {
	// Create a new context with a deadline.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// Add a custom value to the context for logging.
	ctx = context.WithValue(ctx, "requestID", "12345")

	// Start the data fetch operation.
	data, err := simulateDataFetch(ctx, getRequestID(ctx), time.Second*3)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	// Output the fetched data.
	log.Println("Fetched data:", string(data))
}

// getRequestID retrieves the request ID from the context.
func getRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value("requestID").(string)
	if !ok {
		return "no request ID found"
	}
	return requestID
}
