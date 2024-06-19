package main

import (
	"fmt"
)

func main() {
	// Test cases
	fmt.Println(itoa(123))   // Expected output: "123"
	fmt.Println(itoa(-456))  // Expected output: "-456"
	fmt.Println(itoa(0))     // Expected output: "0"
	fmt.Println(itoa(7890))  // Expected output: "7890"
	fmt.Println(itoa(-9876)) // Expected output: "-9876"
}

func itoa(n int) string {
	// Handle the special case of 0
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	// Convert each digit to a character and build the string in reverse
	var result []byte
	for n > 0 {
		digit := n % 10
		result = append(result, byte(digit)+'0')
		n /= 10
	}

	// Reverse the string since we built it backwards
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return sign + string(result)
}
