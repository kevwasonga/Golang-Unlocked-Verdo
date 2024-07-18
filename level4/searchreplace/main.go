package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Retrieve the first argument as the input string
	str := os.Args[1]

	// Convert the second and third arguments to runes (single characters)
	or := rune(os.Args[2][0])
	nr := rune(os.Args[3][0])

	// Replace occurrences of the old rune with the new rune in the input string
	result := replace(str, or, nr)

	// Print the resulting string
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

// replace takes a string and two runes and returns a new string
// where all occurrences of the old rune are replaced with the new rune
func replace(s string, or, nr rune) string {
	var result []rune

	// Iterate over each character in the string
	for _, char := range s {
		// If the character matches the old rune, append the new rune
		if char == or {
			result = append(result, nr)
		} else { // Otherwise, append the original character
			result = append(result, char)
		}
	}

	// Convert the slice of runes back to a string and return it
	return string(result)
}
