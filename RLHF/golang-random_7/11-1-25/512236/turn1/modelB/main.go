package main

import (
	"fmt"
	"strings"
)

// Define a struct to hold user chat engagement metrics
type ChatMetrics struct {
	User     string
	Messages []string
}

// Function to store user chat engagement metrics in slices and analyze the frequency of messages sent per user.
func analyzeChatEngagement(messages []string) (map[string]int, error) {
	// Create an empty map to store user engagement metrics
	userEngagement := make(map[string]*ChatMetrics)

	// Process each message in the input slice
	for _, message := range messages {
		// Split the message into user and text parts using the colon (:) separator
		parts := strings.Split(message, ": ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Invalid message format: %s", message)
		}
		user := parts[0]
		text := parts[1]

		// If the user doesn't exist in the map, add them with an empty metrics struct
		if _, ok := userEngagement[user]; !ok {
			userEngagement[user] = &ChatMetrics{User: user}
		}

		// Append the message text to the user's metrics
		userEngagement[user].Messages = append(userEngagement[user].Messages, text)
	}

	// Create a map to store the frequency of messages sent per user
	messageFrequency := make(map[string]int)

	// Calculate the frequency of messages for each user
	for _, metrics := range userEngagement {
		messageFrequency[metrics.User] = len(metrics.Messages)
	}

	return messageFrequency, nil
}

func main() {
	messages := []string{"alice: Hi everyone!", "bob: Hello Alice!", "chris: Good morning!", "alice: How are you?", "bob: I'm fine, thanks for asking.", "chris: Nice to see you both.", "alice: Bye for now!"}

	frequency, err := analyzeChatEngagement(messages)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print the frequency of messages sent per user
	fmt.Println("Message Frequency per User:")
	for user, freq := range frequency {
		fmt.Printf("%s: %d\n", user, freq)
	}
}
