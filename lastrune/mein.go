package main

import (
	"fmt"
)

func LastRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	return rune(s[len(s)-1])
}

func main() {
	s := "Hello, World!"
	fmt.Println(LastRune(s)) // Output: l
}
