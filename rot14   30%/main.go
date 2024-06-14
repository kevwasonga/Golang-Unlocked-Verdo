package main

import "github.com/01-edu/z01"

func main() {
	input := "HELLO HUMAN"
	rott14(input)
}

func rott14(s string) {
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			z01.PrintRune('a' + (char-'a'+14)%26)
		} else if char >= 'A' && char <= 'Z' {
			z01.PrintRune('A' + (char-'A'+14)%26)
		} else {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
