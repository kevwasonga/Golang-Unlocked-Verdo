package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var result string
	if len(os.Args) > 1 {
		result = (os.Args[len(os.Args)-3])
	}

	for _, printable := range result {
		z01.PrintRune(printable)
	}

	z01.PrintRune('\n')
}
