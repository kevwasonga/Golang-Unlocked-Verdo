package main

import "fmt"

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

// FoldInt function applies the function f to each element in the slice a and accumulates the result
func FoldInt(f func(int, int) int, a []int, n int) {
	result := n
	for _, value := range a {
		result = f(result, value)
	}
	fmt.Println(result)
}

// Add function returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Mul function returns the product of two integers
func Mul(a, b int) int {
	return a * b
}

// Sub function returns the difference of two integers
func Sub(a, b int) int {
	return a - b
}
