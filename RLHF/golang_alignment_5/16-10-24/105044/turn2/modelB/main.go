package main

import (
	"fmt"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

func main() {
	// Sample input text
	text := "The quick brown fox jumps over the lazy dog. The quick brown fox is really fast."

	// Text Compression using Linguistic Methods
	compressedText := compressText(text)

	// Display the compressed text
	fmt.Println("Compressed Text:", compressedText)
}

func compressText(text string) string {
	// Load a Japanese tokenizer
	t := kagome.New()
	defer t.Close()

	// Tokenize the text
	tokens := t.Tokenize(text)

	// Perform NER
	nerTokens := ner(tokens)

	// Perform POS tagging and build phrase dictionary
	phraseDict := buildPhraseDict(nerTokens)

	// Compress the text using phrase dictionary
	compressedText := compressTextWithPhrases(text, phraseDict)

	return compressedText
}

func ner(tokens []tokenizer.Token) []tokenizer.Token {
	// NER simulation using simple heuristics (In real-world scenarios, use dedicated NER tools)
	nerTokens := make([]tokenizer.Token, 0, len(tokens))
	for _, token := range tokens {
		if strings.HasPrefix(token.Surface, "The") {
			nerTokens = append(nerTokens, tokenizer.Token{Surface: "DT", POS: "DT", Lemma: "DT"})
		} else if strings.HasPrefix(token.Surface, "quick") {
			nerTokens = append(nerTokens, tokenizer.Token{Surface: "JJ", POS: "JJ", Lemma: "JJ"})
		} else {
			nerTokens = append(nerTokens, token)
		}
	}
	return nerTokens
}

func buildPhraseDict(tokens []tokenizer.Token) map[string]int {
	phraseDict := make(map[string]int)

	for i := 0; i < len(tokens); i++ {
		for j := i; j < len(tokens); j++ {
			phrase := strings.Join(tokenizer.TokensToString(tokens[i:j+1]), " ")
			phraseDict[phrase]++
		}
	}

	// Remove common words and frequent phrases (can be adjusted based on requirements)
	delete(phraseDict, " ")
	delete(phraseDict, "The")
	delete(phraseDict, "is")
	delete(phraseDict, "quick brown fox")

	return phraseDict
}

func compressTextWithPhrases(text string, phraseDict map[string]int) string {
	// Split the text into phrases
	phrases := strings.Split(text, " ")

	compressedText := make([]string, 0, len(phrases))

	for _, phrase := range phrases {
		if count, ok := phraseDict[phrase]; ok && count > 1 {
			// Compressed phrase: Use a code to represent the frequent phrase
			compressedText = append(compressedText, fmt.Sprintf("%d", len(phraseDict)))
		} else {
			// Retain the original phrase if it's not frequent
			compressedText = append(compressedText, phrase)
		}
	}
	return strings.Join(compressedText, " ")
}
