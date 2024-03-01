package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("résumé")
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe length of the string is %d\n", len(myString))

	var myRune = 'a'
	fmt.Printf("\n myRune = %v", myRune)

	var strSlice = []string{"h", "e", "l", "l", "o", " ", "c", "a", "t"}
	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n catStr = %v", catStr)
}
