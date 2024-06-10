package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	param := len(os.Args) - 1 // Subtract 1 to exclude the program name
	if param < 10 {
		z01.PrintRune('0' + rune(param))
	} else {
		tens := param / 10
		ones := param % 10
		z01.PrintRune('0' + rune(tens))
		z01.PrintRune('0' + rune(ones))
	}
	z01.PrintRune('\n')
}
