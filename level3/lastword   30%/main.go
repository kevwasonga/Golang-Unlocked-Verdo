package main

import "fmt"

func main() {
	input := " hallla madrid"
	fmt.Println(lastword(input))
}

func lastword(s string) string {
	var result string
	j := len(s) - 1

	// Skip trailing spaces
	for j >= 0 && string(s[j]) == " " {
		j--
	}

	// Extract the last word
	for i := j; i >= 0 && string(s[i]) != " "; i-- {
		result = string(s[i]) + result
	}

	return result
}
