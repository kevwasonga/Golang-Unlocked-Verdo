package main

import "github.com/01-edu/z01"

func main() {
	input := "hello man"

	rot13(input)
}

func rot13(s string) {
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			z01.PrintRune('a' + (char-'a'+13)%26)
		} else if char >= 'A' && char <= 'Z' {
			z01.PrintRune(char + (char-'A'+13)%26)
		} else {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
