package main

import (
	"fmt"
)

func main() {
	fmt.Println(Unzipstring("2f7h2a"))   // Output:
	fmt.Println(Unzipstring("2O5u2H0e")) // Output: OuuuuuHH
}

// Unzipstring function takes a coded string and returns its decoded form
func Unzipstring(s string) string {
	lenghth := len(s)

	var result string

	if len(s) == 0 {
		return "Invalid input "
	}

	for i := 0; i < len(s); i += 2 {
		if i+1 > lenghth || !isDigit(rune((s[i]))) || !isLetter(rune((s[i+1]))) {
			return "Invalid input "
		}

		count := int(s[i] - '0')
		char := string(s[i+1])

		for j := 0; j < count; j++ {
			result += char
		}
	}
	return result
}

func isDigit(c rune) bool {
	return (c >= '0' && c <= '9')
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
