package main

import (
	"fmt"
	"sort"
)

// Message represents a single message sent by a user
type Message struct {
	User    string
	Content string
}

// Messages stores a collection of messages
type Messages []Message

// AddMessage appends a new message to the Messages collection
func (msgs *Messages) AddMessage(user, content string) {
	*msgs = append(*msgs, Message{User: user, Content: content})
}

// AnalyzeMessageFrequency calculates the frequency of messages sent by each user
func (msgs Messages) AnalyzeMessageFrequency() map[string]int {
	userCount := make(map[string]int)

	for _, msg := range msgs {
		userCount[msg.User]++
	}

	return userCount
}

// PrintMessageFrequency prints the message frequency in a sorted manner
func PrintMessageFrequency(userCount map[string]int) {
	users := make([]string, 0, len(userCount))
	for user := range userCount {
		users = append(users, user)
	}
	sort.Strings(users)

	for _, user := range users {
		fmt.Printf("%s: %d messages\n", user, userCount[user])
	}
}

func main() {
	// Example usage
	messages := Messages{}

	messages.AddMessage("Alice", "Hello!")
	messages.AddMessage("Bob", "Hi there!")
	messages.AddMessage("Alice", "How are you?")
	messages.AddMessage("Bob", "I'm doing well, thanks.")
	messages.AddMessage("Charlie", "Good day!")
	messages.AddMessage("Alice", "Same here!")

	userCount := messages.AnalyzeMessageFrequency()
	PrintMessageFrequency(userCount)
}
