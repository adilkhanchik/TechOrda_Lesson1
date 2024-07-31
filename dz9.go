package main

import (
	"fmt"
	"strconv"
)

func DecimalToBinary(decimal int) string {
	return strconv.FormatInt(int64(decimal), 2)
}

func main() {
	decimal := 13
	binary := DecimalToBinary(decimal)
	fmt.Printf("The binary representation of %d is %s\n", decimal, binary)
}
