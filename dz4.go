package main

import (
	"fmt"
)

// FindWord searches for a lowercase word in a string of uppercase letters.
func FindWord(crowd string, word string) string {
	wordLen := len(word)
	crowdLen := len(crowd)

	if wordLen == 0 {
		return ""
	}

	if crowdLen == 0 {
		return ""
	}

	// Initialize pointers for the word and crowd strings
	wordIndex := 0

	// Traverse the uppercase string
	for i := 0; i < crowdLen; i++ {
		if wordIndex < wordLen && crowd[i] == word[wordIndex] {
			wordIndex++
			if wordIndex == wordLen {
				return word
			}
		}
	}

	// If we didn't find the whole word
	return ""
}

func main() {
	// Example usage
	crowd := "XABYCDZEFGHIJKLMNOPQRS"
	word := "abc"
	result := FindWord(crowd, word)
	if result != "" {
		fmt.Printf("The word '%s' was found in the crowd.\n", result)
	} else {
		fmt.Println("The word was not found in the crowd.")
	}
}
