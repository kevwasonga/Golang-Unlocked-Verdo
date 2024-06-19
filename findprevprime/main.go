package main

import "fmt"

func main() {
	// Print the result of FindPrevPrime for different inputs
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(12))
}

// FindPrevPrime finds the largest prime number less than or equal to nb
func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}

	for i := 2; i < nb/2; {
		if nb%2 == 0 {
			nb--
			i = 2
		} else {
			i++
		}
	}
	return nb

}
