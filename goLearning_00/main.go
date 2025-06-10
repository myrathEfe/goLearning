package main

import (
	"errors"
	"fmt"
)

func main() {

	var printThis = "Hello World!"
	printMe(printThis)

	num1, num2 := 6, 4

	result, remainder, err := intDivision(num1, num2)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is : %v\n", result)
	} else {
		fmt.Printf("The result of the integer division is : %v with remainder %v", result, remainder)
	}
}

func printMe(key string) {
	fmt.Println(key)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("denominator is zero")
		return 0, 0, err
	}
	var result = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder, err
}
