package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Need to use
	var intNum int64
	intNum += 1

	var floatNum float64
	floatNum = 3.14

	var myString string
	myString = "Yikuanzz"

	var myBool bool
	myBool = true

	fmt.Printf("int number: %d\n", intNum)
	fmt.Printf("float number: %.2f\n", floatNum)
	fmt.Printf("my string: %s\n", myString)
	fmt.Println(utf8.RuneCountInString("Y"))
	fmt.Printf("myBool: %t\n", myBool)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value"
	const pi float64 = 3.1
	fmt.Println(myConst, pi)
}
