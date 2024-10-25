package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/clips/pattern"
	"github.com/siddontang/go-log/logger"
	"github.com/topictoolkit/lda"
)

// LogEntry represents a structured log entry
type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Level     string `json:"level"`
	UserID    int    `json:"user_id,omitempty"`
	ErrorCode int    `json:"error_code,omitempty"`
}

func main() {
	logMessages := []string{
		`{"timestamp": "2023-07-28T10:00:01Z", "message": "User logged in successfully [user_id=456]", "level": "info"}`,
		`{"timestamp": "2023-07-28T10:00:02Z", "message": "Order placed successfully [user_id=123, order_id=789]", "level": "info"}`,
		`{"timestamp": "2023-07-28T10:00:03Z", "message": "System error: Internal Server Error (500)", "level": "error"}`,
		`{"timestamp": "2023-07-28T10:00:04Z", "message": "Product stock updated [product_id=abc]", "level": "info"}`,
		`{"timestamp": "2023-07-28T10:00:05Z", "message": "User feedback received: Great app!", "level": "info"}`,
	}

	// 1. Structured Log Parsing
	var entries []LogEntry
	for _, msg := range logMessages {
		var entry LogEntry
		if err := json.Unmarshal([]byte(msg), &entry); err != nil {
			log.Printf("Error parsing log entry: %v", err)
			continue
		}
		entries = append(entries, entry)
	}

	// 2. Sentiment Analysis
	sentimentAnalyzer := pattern.NewSentimentAnalyzer(pattern.English)
	for _, entry := range entries {
		sentiment := sentimentAnalyzer.Sentiment(strings.ToLower(entry.Message))
		entry.Message += fmt.Sprintf(" (Sentiment: %s)", sentiment)
	}

	// 3. Text Preprocessing and Template Matching
	template := "{{timestamp}} {{level}}: {{message}} [User ID: {{user_id}}]"
	for i, entry := range entries {
		preprocessedMessage := strings.Map(func(r rune) rune {
			if r == '.' || r == ',' || r == '!' || r == '?' {
				return -1
			}
			return r
		}, strings.ToLower(entry.Message))

		rendered := raymond.Render(template, map[string]interface{}{
			"timestamp": entry.Timestamp,
			"level":     entry.Level,
			"message":   preprocessedMessage,
			"user_id":   entry.UserID,
		})
		entries[i].Message = rendered
	}

	// 4. Topic Modeling (Using a simplified example for demonstration)
	documents := make([][]string, len(entries))
	for i, entry := range entries {
		documents[i] = strings.Fields(strings.ToLower(entry.Message))
	}

	numTopics := 2
	model, err := lda.NewLDA(numTopics, documents, 50, 0.01, 0.01)
	if err != nil {
		logger.Fatalf("Error running LDA: %v", err)

	}
}
