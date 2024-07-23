package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countConsecutiveChars("eeoe"))        // Output: 2e1o1e
	fmt.Println(countConsecutiveChars("hello"))       // Output: 1h1e2l1o
	fmt.Println(countConsecutiveChars("mississippi")) // Output: 1m1i2s1i2s1i2p1i
	fmt.Println(countConsecutiveChars("aabbbccaa"))   // Output: 2a3b2c2a
	fmt.Println(countConsecutiveChars("aabbaaccaol"))
}

// countConsecutiveChars counts the consecutive occurrences of characters in the input string
// and produces a string showing the count followed by the character.
func countConsecutiveChars(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result string
	prev := rune(input[0])
	count := 1

	for i := 1; i < len(input); i++ {
		char := rune(input[i])
		if char == prev {
			count++
		} else {
			result += strconv.Itoa(count) + string(prev)
			count = 1
			prev = char
		}
	}
	// Append the last set of characters
	result += strconv.Itoa(count) + string(prev)

	return result
}
