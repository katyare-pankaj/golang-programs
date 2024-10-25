package main

import (
	"fmt"
	"strings"

	"github.com/siddontang/go-log/log"
	"github.com/topictoolkit/lda"
)

func main() {
	logMessages := []string{
		"User logged in",
		"Order placed successfully",
		"System performance issue detected",
		"User feedback received",
		"Product stock updated",
	}
	// Create a document-term matrix
	documents := make([][]string, len(logMessages))
	for i, msg := range logMessages {
		documents[i] = strings.Fields(strings.ToLower(msg))
	}
	// Run LDA topic modeling
	numTopics := 2
	model, err := lda.NewLDA(numTopics, documents, 50, 0.01, 0.01)
	if err != nil {
		log.Fatalf("Error running LDA: %v", err)
	}
	// Display the top words for each topic
	for topicIndex, topic := range model.Topics {
		fmt.Printf("Topic %d: %v\n", topicIndex, topic.TopWords(10))
	}
}
