package main

import (
	"fmt"
)

// sumUpTo takes an integer n and returns the sum of all numbers from 1 to n.
func sumUpTo(n int) int {
	// Initialize a variable to hold the sum
	sum := 0

	// Loop from 1 to n, adding each number to sum
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

func main() {
	// Test the function with an example input
	result := sumUpTo(4)
	fmt.Println(result) // Output: 10
}
