package main

import "fmt"

func main() {
	input := "hello"

	fmt.Println(FirstRune(input)) // Call FirstRune with the input string
}

func FirstRune(s string) rune {
	var result rune

	for _, char := range s { // Iterate over each rune in the string
		result = char // Assign the first rune to result
		break         // Exit the loop after the first rune
	}
	return result // Return the first rune
}
