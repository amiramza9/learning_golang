package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	pl := fmt.Println
	word := "amir hamza"
	pl("Rune Count: ", utf8.RuneCountInString(word))
	for i, runeval := range word {
		fmt.Printf("%d : %#U : %c\n", i, runeval, runeval)
	}
}
