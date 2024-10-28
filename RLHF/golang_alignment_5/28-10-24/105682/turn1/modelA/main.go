package main

import (
	"fmt"
	"strings"
)

// A Token is a simple struct to represent a token in NLP
type Token struct {
	Text  string
	Tag   string
	Lemma string
}

// A Sentence is a slice of Tokens
type Sentence []Token

// filterTokens is a higher-order function that takes a predicate function
// as an argument and returns a new function that filters a Sentence based on that predicate.
func filterTokens(predicate func(Token) bool) func(Sentence) Sentence {
	return func(s Sentence) Sentence {
		var filtered Sentence
		for _, t := range s {
			if predicate(t) {
				filtered = append(filtered, t)
			}
		}
		return filtered
	}
}

// isNoun is a predicate function to check if a Token is a noun
func isNoun(t Token) bool {
	return t.Tag == "NN"
}

// isVerb is a predicate function to check if a Token is a verb
func isVerb(t Token) bool {
	return t.Tag == "VB"
}

// lemmatizeTokens is a higher-order function that applies a lemmatization function
// to each Token in a Sentence.
func lemmatizeTokens(lemmatizer func(string) string) func(Sentence) Sentence {
	return func(s Sentence) Sentence {
		lemmatized := make(Sentence, len(s))
		for i, t := range s {
			lemmatized[i] = Token{Text: t.Text, Tag: t.Tag, Lemma: lemmatizer(t.Text)}
		}
		return lemmatized
	}
}

// simpleLemmatizer is a basic lemmatization function for demonstration purposes.
func simpleLemmatizer(word string) string {
	return strings.ToLower(word)
}

func main() {
	sentence := Sentence{
		{Text: "The", Tag: "DT", Lemma: "the"},
		{Text: "quick", Tag: "JJ", Lemma: "quick"},
		{Text: "brown", Tag: "JJ", Lemma: "brown"},
		{Text: "fox", Tag: "NN", Lemma: "fox"},
		{Text: "jumps", Tag: "VB", Lemma: "jump"},
		{Text: "over", Tag: "IN", Lemma: "over"},
		{Text: "the", Tag: "DT", Lemma: "the"},
		{Text: "lazy", Tag: "JJ", Lemma: "lazy"},
		{Text: "dog", Tag: "NN", Lemma: "dog"},
	}

	// Using partial functions to create filtered functions
	filterNouns := filterTokens(isNoun)
	filterVerbs := filterTokens(isVerb)

	// Using composition to create a pipeline
	processSentence := lemmatizeTokens(simpleLemmatizer)
	filteredNouns := processSentence(filterNouns(sentence))
	filteredVerbs := processSentence(filterVerbs(sentence))

	fmt.Println("Original Sentence:")
	fmt.Println(sentence)

	fmt.Println("\nNouns after lemmatization:")
	fmt.Println(filteredNouns)

	fmt.Println("\nVerbs after lemmatization:")
	fmt.Println(filteredVerbs)
}
