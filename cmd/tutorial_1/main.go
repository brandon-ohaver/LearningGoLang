package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe length of 'myString' is %v\n", len(myString))

	var myRune = 'a'
	fmt.Printf("myRune = %v\n", myRune)

	var strSlice = []string {"b", "r", "a", "n", "d" , "o", "n"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("%v\n", catStr)
}
