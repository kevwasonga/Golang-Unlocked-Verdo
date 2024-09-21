package main

import "fmt"

func main() {
	input := "MCMXCIV" // Test input for Roman numeral

	fmt.Println(romans(input)) // Convert and print the result
}

// Function to convert Roman numeral string to integer
func romans(s string) int {
	// Map of Roman numeral characters to their integer values
	maps := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	previous := 0 // Variable to hold the previous numeral value
	var total int // Variable to accumulate the total value

	// Iterate through the string from the last character to the first
	for i := len(s) - 1; i >= 0; i-- {
		current := maps[s[i]] // Get the current numeral value

		// If the previous numeral is greater, we subtract the current value
		if previous > current {
			total -= current
		} else {
			// Otherwise, we add the current value
			total += current
		}
		// Update the previous value to the current
		previous = current
	}

	return total // Return the final calculated total
}
