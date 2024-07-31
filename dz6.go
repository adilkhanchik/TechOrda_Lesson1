package main

import (
	"fmt"
	"strings"
)

func ReplaceWordsWithEmojis(sentence string) string {
	emojiMap := map[string]string{
		"smile": ":)",
		"grin":  ":D",
		"sad":   ":(",
		"mad":   ">_<",
	}

	words := strings.Fields(sentence)

	for i, word := range words {
		if emoji, found := emojiMap[strings.ToLower(word)]; found {
			words[i] = emoji
		}
	}

	return strings.Join(words, " ")
}

func main() {
	sentence := "I feel so smile today but sometimes I grin and then I get sad or even mad."
	result := ReplaceWordsWithEmojis(sentence)
	fmt.Println(result)
}
