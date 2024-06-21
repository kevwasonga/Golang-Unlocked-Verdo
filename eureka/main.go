package main

import (
	"fmt"
	"math"

	"strconv"
)

func main() {
	// Define the range for testing
	a, b := uint64(1), uint64(1000)

	// Call SumDigPow to get the numbers that satisfy the condition
	result := SumDigPow(a, b)

	// Print the result
	fmt.Printf("Numbers between %d and %d that satisfy the 'eureka' condition:\n", a, b)
	for _, num := range result {
		fmt.Println(num)
	}
}

func SumDigPow(a, b uint64) []uint64 {
	var result []uint64

	for val := a; val <= b; val++ {
		if isEureka(val) {
			result = append(result, val)
		}
	}

	return result
}

func isEureka(n uint64) bool {
	var result uint64
	val := strconv.FormatUint(n, 10)

	for i, char := range val {

		pos := i + 1
		dig, _ := strconv.Atoi(string(char))

		result += uint64(math.Pow(float64(dig), float64(pos)))

	}
	return result == n
}
