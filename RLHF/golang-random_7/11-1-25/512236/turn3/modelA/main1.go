package main

import (
	"fmt"
	"sort"
	"strings"
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

// AnalyzeWordFrequency calculates the number of words for each user, excluding empty or space-only messages
func (msgs Messages) AnalyzeWordFrequency() map[string]int {
	wordCount := make(map[string]int)

	for _, msg := range msgs {
		trimmedContent := strings.TrimSpace(msg.Content)
		// Skip empty or space-only messages
		if trimmedContent == "" {
			continue
		}
		words := strings.Fields(trimmedContent)
		// Skip messages with only one word
		if len(words) > 1 {
			wordCount[msg.User] += len(words) // Increment the count by the number of words
		}
	}

	return wordCount
}

// PrintWordFrequency prints the word frequency in a sorted manner
func PrintWordFrequency(userCount map[string]int) {
	users := make([]string, 0, len(userCount))
	for user := range userCount {
		users = append(users, user)
	}
	sort.Strings(users)

	for _, user := range users {
		fmt.Printf("%s: %d words\n", user, userCount[user])
	}
}

func main() {
	// Example usage
	messages := Messages{}

	messages.AddMessage("Alice", "Hello!")
	messages.AddMessage("Bob", " ")
	messages.AddMessage("Alice", "Good morning!")
	messages.AddMessage("Bob", "Fine, thanks.")
	messages.AddMessage("Charlie", "   ")
	messages.AddMessage("Alice", "Sure, let me know.")

	wordCount := messages.AnalyzeWordFrequency()
	PrintWordFrequency(wordCount)
}
