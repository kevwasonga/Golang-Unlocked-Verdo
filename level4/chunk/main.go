package main

import (
	"github.com/01-edu/z01"
)

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	if size == 0 {
		z01.PrintRune('\n')
		return
	}

	var result [][]int
	var chunk []int

	for _, val := range slice {
		chunk = append(chunk, val)
		if len(chunk) == size {
			result = append(result, chunk)
			chunk = []int{}
		}
	}

	// Append remaining elements if any
	if len(chunk) > 0 {
		result = append(result, chunk)
	}

	// Print the chunks
	z01.PrintRune('[')

	for _, val := range result {
		z01.PrintRune('[')

		for _, char := range val {
			z01.PrintRune(rune(char + '0'))
			z01.PrintRune(' ')
		}

		z01.PrintRune(']')

	}
	z01.PrintRune(']')
	z01.PrintRune('\n')
}
