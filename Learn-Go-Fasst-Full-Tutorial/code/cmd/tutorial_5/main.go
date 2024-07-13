package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "Yikuanzz"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\n The length of String is %d\n", len(myString))

	var strBuilder strings.Builder
	var strSlice = []string{"s", "u", "b", "s", "c"}
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)

}
