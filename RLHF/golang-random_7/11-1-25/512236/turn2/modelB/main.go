package main

import (
	"strings"
)

// AnalyzeMessageFrequency calculates the frequency of messages sent by each user, excluding messages with only one word.
func (msgs Messages) AnalyzeMessageFrequency() map[string]int {
	userWordCount := make(map[string]int)

	for _, msg := range msgs {
		words := strings.Fields(msg.Content)
		if len(words) > 1 { // Exclude messages with only one word
			userWordCount[msg.User] += len(words)
		}
	}

	return userWordCount
}
