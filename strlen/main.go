package main

import "fmt"

func main() {
	l := ("hello Verdo")
	fmt.Println(StrLen(l))
}

func StrLen(s string) int {
	var count int

	for range s {
		count++
	}

	return count
}
