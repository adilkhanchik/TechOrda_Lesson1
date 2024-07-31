package main

import (
	"fmt"
)

func CountSquares(n int) int {
	total := 0
	for k := 1; k <= n; k++ {
		total += (n - k + 1) * (n - k + 1)
	}
	return total
}

func main() {
	n := 3
	result := CountSquares(n)
	fmt.Printf("Number of different squares in a %d x %d grid: %d\n", n, n, result)
}
