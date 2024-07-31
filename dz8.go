package main

import (
	"fmt"
	"strings"
)

// CountVowels counts the number of vowels in a given string.
func CountVowels(input string) int {
	// Define a set of vowels (both lowercase and uppercase)
	vowels := "aeiouAEIOU"

	// Initialize a counter for vowels
	count := 0

	// Iterate through each character in the input string
	for _, char := range input {
		// Check if the character is a vowel
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}

func main() {
	// Example usage
	input := "Hello, World!"
	vowelCount := CountVowels(input)
	fmt.Printf("The number of vowels in '%s' is %d\n", input, vowelCount)
}
