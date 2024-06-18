package main

import "fmt"

func main() {
	input := "-123" // Example input string

	fmt.Println(atoi(input)) // Call atoi function and print the result
}

func atoi(s string) int {
	var result int // Initialize result to store the converted integer
	sign := 1      // Initialize sign to 1 (positive)

	// Check for negative sign at the beginning of the string
	if s[0] != ' ' && s[0] == '-' {
		sign = -1 // Update sign to -1 for negative numbers
		s = s[1:] // Remove the negative sign from the string
	}

	// Iterate over each character in the string
	for _, char := range s {
		if char < '0' || char > '9' { // Check if the character is not a digit
			return 0 // Return 0 for invalid characters
		}
		// Convert character to integer and update result
		result = result*10 + int(char-'0')
	}

	return result * sign // Multiply result by sign to account for negative numbers
}
