package main

import (
	"fmt"
)

func sumUpTo(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

func main() {
	result := sumUpTo(4)
	fmt.Println(result)
}
