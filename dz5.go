package main

import (
	"fmt"
	"strings"
)

func CountPotatoes(input string) int {
	lowercaseInput := strings.ToLower(input)

	word := "potato"

	count := 0
	for {
		index := strings.Index(lowercaseInput, word)
		if index == -1 {
			break
		}

		count++

		lowercaseInput = lowercaseInput[index+len(word):]
	}

	return count
}

func main() {
	input := "Potato soup is great. I love potatoes. The potato salad is also good. Potato!"
	count := CountPotatoes(input)
	fmt.Printf("The word 'potato' appears %d times in the string.\n", count)
}
