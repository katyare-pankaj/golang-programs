package main

import (
	"fmt"
	"strings"

	"github.com/jdkato/prose/v2"
)

func posTaggingAndCompression(text string) string {
	doc := prose.NewDocument(text)
	tokens := doc.Tokens()
	compressedText := []string{}

	for _, token := range tokens {
		pos := token.Tag
		// Perform compression based on POS tags
		switch pos {
		case "DT": // Determiner
			continue // Skip articles
		default:
			compressedText = append(compressedText, token.Text)
		}
	}
	return strings.Join(compressedText, " ")
}

func namedEntityRecognitionAndCompression(text string) string {
	doc := prose.NewDocument(text)
	entities := doc.Entities()
	compressedText := []string{}
	seenEntities := make(map[string]bool)

	for _, ent := range entities {
		if !seenEntities[ent.Text] {
			compressedText = append(compressedText, ent.Text)
			seenEntities[ent.Text] = true
		}
	}
	return strings.Join(compressedText, " ")
}

func stemmingCompression(text string) string {
	t := tokenizer.New()
	tokens := t.Tokenize(text)
	stemmer := stem.NewPorterStemmer()
	compressedText := []string{}

	for _, token := range tokens {
		stemmed := stemmer.Stem(token.Text())
		compressedText = append(compressedText, stemmed)
	}
	return strings.Join(compressedText, " ")
}

func main() {
	sampleText := `
	Welcome to our exciting new website! We are thrilled to launch it and provide you with the best online shopping experience.
	Shop now and discover a wide range of products at unbeatable prices. Our team of experts is dedicated to ensuring you find the perfect item for your needs.
	`

	fmt.Println("Original Text:")
	fmt.Println(sampleText)

	fmt.Println("\nText after POS Tagging and Compression:")
	compressedText := posTaggingAndCompression(sampleText)
	fmt.Println(compressedText)

	fmt.Println("\nText after Named Entity Recognition and Compression:")
	compressedText = namedEntityRecognitionAndCompression(sampleText)
	fmt.Println(compressedText)

	fmt.Println("\nText after Stemming Compression:")
	compressedText = stemmingCompression(sampleText)
	fmt.Println(compressedText)
}
