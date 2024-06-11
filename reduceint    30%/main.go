package main

import "fmt"

// The main function where the program execution begins
func main() {
	// Define a multiplication function
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	// Define an addition function
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	// Define a division function
	div := func(acc int, cur int) int {
		return acc / cur
	}

	// Create a slice of integers
	as := []int{500, 2}

	// Call ReduceInt with the slice and the multiplication function
	ReduceInt(as, mul) // Should print 1000
	// Call ReduceInt with the slice and the addition function
	ReduceInt(as, sum) // Should print 502
	// Call ReduceInt with the slice and the division function
	ReduceInt(as, div) // Should print 250
}

// ReduceInt applies a function to each element of the slice and prints the result
func ReduceInt(a []int, f func(int, int) int) {
	// Initialize the accumulator with the first element of the slice
	acc := a[0]

	// Iterate over the slice starting from the second element
	for _, val := range a[1:] {
		// Apply the function f to the accumulator and the current element
		acc = f(acc, val)
	}

	// Print the final accumulated result
	fmt.Println(acc)
}
