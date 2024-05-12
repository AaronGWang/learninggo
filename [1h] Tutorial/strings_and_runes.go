package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("résumé")
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString{
		fmt.Printf("Index: %v | Value: %v\n", i, v)
	}

	var myRune = "a"
	fmt.Printf("%v\n", myRune)

	var strSlice = []string{"t", "e", "s", "t"}
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("%v\n", catStr)
}