package main

import "fmt"

func main() {
	input := "MCMXCIV" //"XXII"

	fmt.Println(romans(input))
}

func romans(s string) int {
	maps := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	previous := 0

	var total int

	for i := len(s) - 1; i >= 0; i-- {
		current := maps[s[i]]

		if previous > current {
			total -= current
		} else {
			total += current
		}
		//update the last processed to be the new "previous"
		previous = current

	}

	return total
}
