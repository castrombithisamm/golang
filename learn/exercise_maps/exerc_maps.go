package main

import (
	"strings"
	// "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// Initialize an empty map to store word counts
	wordCounts := make(map[string]int)

	// Split the input string into words
	words := strings.Fields(s)

	// Iterate over the words and update the counts
	for _, word := range words {
		// Increment the count for the current word
		wordCounts[word]++
	}

	return wordCounts
}

func main() {
	// Test the WordCount function
	// wc.Test(WordCount)
}
