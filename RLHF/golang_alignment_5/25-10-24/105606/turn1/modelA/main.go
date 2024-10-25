package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nlpodyssey/spago/pkg/ner"
	"github.com/nlpodyssey/spago/pkg/pos"
	"github.com/nlpodyssey/spago/pkg/tokenizers/sentencepiece"
)

func main() {
	logMessages := []string{
		"Server started on port 8080",
		"User JohnDoe logged in successfully",
		"Product ABC was added to cart",
		"Error processing order #123: invalid quantity",
	}

	// Regular Expression Example
	for _, msg := range logMessages {
		if matched, _ := regexp.MatchString(`Error.*`, msg); matched {
			fmt.Println("Error Log:", msg)
		}
	}

	// Named Entity Recognition (NER) Example
	model, err := ner.LoadModel("en")
	if err != nil {
		panic(err)
	}
	tokenizer := sentencepiece.NewTokenizer(model.Vocab())

	for _, msg := range logMessages {
		tokens := tokenizer.EncodeAsIds(msg)
		tags := model.Predict(tokens)
		fmt.Println("NER Analysis for:", msg)
		for _, ent := range ner.Entities(msg, tokens, tags) {
			fmt.Printf("-%s: %s\n", ent.Label, ent.Text)
		}
		fmt.Println()
	}

	// Part-of-Speech (POS) Tagging Example
	posModel, err := pos.LoadModel("en")
	if err != nil {
		panic(err)
	}

	for _, msg := range logMessages {
		tokens := tokenizer.EncodeAsIds(msg)
		tags := posModel.Predict(tokens)
		fmt.Println("POS Tagging for:", msg)
		for _, tok := range pos.Tag(msg, tokens, tags) {
			fmt.Printf("-%s: %s\n", tok.Text, tok.Tag)
		}
		fmt.Println()
	}

	// Text Cleaning Example
	stopWords := []string{"the", "is", "and", "a", "an"}
	for _, msg := range logMessages {
		cleanedMsg := cleanText(msg, stopWords)
		fmt.Println("Cleaned Log:", cleanedMsg)
	}
}

func cleanText(text string, stopWords []string) string {
	text = strings.ToLower(text)
	text = regexp.MustCompile(`[^a-zA-Z\s]`).ReplaceAllString(text, "")
	words := strings.Fields(text)
	filteredWords := make([]string, 0)
	for _, word := range words {
		if !contains(stopWords, word) {
			filteredWords = append(filteredWords, word)
		}
	}
	return strings.Join(filteredWords, " ")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
