package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"
	// char := 'o'

	// Convert the character to a string
	// charStr := string(char)
	charStr := "lo"
	strA := "ABC"
	strB := "CBA"
	resul := strings.Compare(strA, strB)
	if resul == 0 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are not equal")
	}

	// Use strings.Contains() to check if the character exists in the string
	if strings.Contains(str, charStr) {
		fmt.Printf("The character '%s' is present in the string.\n", charStr)
	} else {
		fmt.Printf("The character '%s' is not present in the string.\n", charStr)
	}
}
