package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the number of arguments is correct
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	// Get the input string from the command-line argument
	input := os.Args[1]

	// Initialize an empty string to store the result
	var result string

	// Iterate over each character in the input string
	for i := range input {
		// Check if the current character is a letter
		if (input[i] >= 'a' && input[i] <= 'z') || (input[i] >= 'A' && input[i] <= 'Z') {
			// Replace the letter with its opposite
			if input[i] >= 'a' && input[i] <= 'z' {
				result += string('a' + 'z' - input[i])
			} else {
				result += string('A' + 'Z' - input[i])
			}
		} else {
			// For non-letter characters, add them to the result as-is
			result += string(input[i])
		}
	}

	// Print the transformed string
	fmt.Println(result)
}
