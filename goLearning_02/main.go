package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("************************")
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("é karakterinin Unicode kod noktası (decimal): %v, tipi: %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("The length of %q is %v\n", string(myString), len(myString))
	fmt.Println("************************")
	var myString2 = "résumé"
	for i, v := range myString2 {
		fmt.Println(i, v)
	}
	fmt.Printf("The length of %q is %v\n", myString2, len(myString2))
	fmt.Println("************************")

	var r rune = '€' /*Tek tırnak ile tutulur*/
	fmt.Printf("%c - %U - %d\n", r, r, r)
	fmt.Println("************************")

	var strSlice = []string{"m", "e", "r", "h", "a", "b", "a"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("%v", catStr)
}
