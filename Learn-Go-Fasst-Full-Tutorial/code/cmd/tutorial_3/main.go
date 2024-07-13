package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Hello World!")

	numerator, denominator := 11, 3
	result, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("The result of the interger division is %v with remainder %v.", result, remainder)
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return numerator, denominator, err
	}

	var result = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder, nil
}
