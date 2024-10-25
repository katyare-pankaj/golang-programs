package main

import (
	"fmt"
	"strings"

	"github.com/clips/pattern"
)

func main() {
	logMessage := "The product is of excellent quality!"
	sentimentAnalyzer := pattern.NewSentimentAnalyzer(pattern.English)
	sentiment := sentimentAnalyzer.Sentiment(strings.ToLower(logMessage))
	fmt.Println("Sentiment:", sentiment)
}
