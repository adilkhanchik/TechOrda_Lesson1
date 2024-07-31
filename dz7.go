package main

import (
	"fmt"
	"strconv"
)

func HighestDigit(number int) int {
	numStr := strconv.Itoa(number)

	highest := 0

	for _, ch := range numStr {
		digit, err := strconv.Atoi(string(ch))
		if err != nil {
			continue
		}

		if digit > highest {
			highest = digit
		}
	}

	return highest
}

func main() {
	number := 93847
	result := HighestDigit(number)
	fmt.Printf("The highest digit in %d is %d\n", number, result)
}
