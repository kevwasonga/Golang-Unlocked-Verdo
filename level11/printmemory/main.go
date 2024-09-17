package main

import "fmt"

func main() {
	printMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func printMemory(arr [10]byte) {
	var result string
	for _, char := range arr {
		result += btoHex(char) + " "
	}
	fmt.Println(result)

	var asci string
	s := string(arr[:])
	for _, char := range s {
		if char < 32 || char > 126 {
			asci += "."
		} else {
			asci += string(char)
		}
	}

	fmt.Println(asci)
}

func btoHex(b byte) string {
	return string([]byte{(getHex((b >> 4) & 0x0F)), getHex(b & 0x0F)})

	// return output
}

func getHex(b byte) byte {
	if b < 10 {
		return '0' + b
	} else {
		return 'a' + (b - 10)
	}
}
