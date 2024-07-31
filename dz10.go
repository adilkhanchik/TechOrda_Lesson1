package main

import (
	"fmt"
	"strconv"
)

func SumStrings(str1, str2 string) (int, error) {
	num1, err1 := strconv.Atoi(str1)
	if err1 != nil {
		return 0, fmt.Errorf("string: %s cannot be converted", str1)
	}

	num2, err2 := strconv.Atoi(str2)
	if err2 != nil {
		return 0, fmt.Errorf("string: %s cannot be converted", str2)
	}

	return num1 + num2, nil
}

func main() {
	str1 := "123"
	str2 := "456"

	sum, err := SumStrings(str1, str2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The sum of %s and %s is %d\n", str1, str2, sum)
	}

}
