package main

import (
	"fmt"
	"sort"
	"strings"
)

// Function to count word frequencies in a given text
func wordFrequency(text string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		wordCount[lowerWord]++
	}

	return wordCount
}

// Function to compress text by selecting the most frequent words
func compressText(text string, topN int) string {
	frequencies := wordFrequency(text)
	type wordFreq struct {
		word  string
		count int
	}

	// Create a slice of word frequencies
	var freqList []wordFreq
	for word, count := range frequencies {
		freqList = append(freqList, wordFreq{word, count})
	}

	// Sort the slice by frequency
	sort.Slice(freqList, func(i, j int) bool {
		return freqList[i].count > freqList[j].count
	})

	// Build the compressed text from the top N frequent words
	var compressedText []string
	for i := 0; i < topN && i < len(freqList); i++ {
		compressedText = append(compressedText, freqList[i].word)
	}

	return strings.Join(compressedText, " ")
}

func main() {
	// New example text for compression related to social media strategies
	text := "In today's digital age, effective social media strategies are crucial for business growth. Engaging content, timely responses, and active community management play key roles. Brands must analyze performance metrics to refine their approach."

	// Compress the text to get the top 5 frequent words
	compressedText := compressText(text, 5)
	fmt.Println("Original Text:", text)
	fmt.Println("Compressed Text:", compressedText)
}
