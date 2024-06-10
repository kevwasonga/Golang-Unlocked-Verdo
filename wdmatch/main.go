package main

import (
	"fmt"
	"os"
)

func main() {

	test:= os.Args
	input := test[1]
	input2 := test[2]
	var result string
	i := 0
	j := 0
	for i < len(input) && j < len(input2) {
		if string(input[i]) == string(input2[j]) {
			result += string(input[i])
			i++
			j++
		} else {
			j++
		}
	}
	fmt.Println(result)
}
