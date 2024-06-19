package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	if size == 0 {
		return
	}
	var result [][]int
	var Chunk []int
	for _, val := range slice {
		Chunk = append(Chunk, val)

		if len(Chunk) == size {
			result = append(result, Chunk)
			Chunk = []int{}

		}
	}
	//Append the last chunk if it's not empty
	if len(Chunk) > 0 {
		result = append(result, Chunk)

	}

	fmt.Println(result)
}
