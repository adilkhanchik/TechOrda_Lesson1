package main

import (
	"fmt"
	"strconv"
)

// DecimalToBinary converts a base-10 string number to its base-2 (binary) representation.
func DecimalToBinary(decimalStr string) (string, error) {
	// Parse the decimal string to an integer
	decimalInt, err := strconv.Atoi(decimalStr)
	if err != nil {
		return "", err // Return an empty string and the error if parsing fails
	}

	// Convert the integer to a binary string
	binaryStr := strconv.FormatInt(int64(decimalInt), 2)

	return binaryStr, nil
}

func main() {
	// Test the function with an example input
	decimalStr := "10"
	binaryStr, err := DecimalToBinary(decimalStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Decimal %s in binary is %s\n", decimalStr, binaryStr)
}
