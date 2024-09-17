package main

import "fmt"

func main() {
	input := byte(254)
	input1 := 5156454

	fmt.Println(byToHEX(input))
	fmt.Println(decToHex(input1))
}

func getHex(b byte) byte {
	if b < 10 {
		return '0' + b
	} else {
		return 'a' + b - 10
	}
}

func decToHex(n int) string {
	hex := ""

	for n > 0 {
		remainder := n % 16

		hex = string(getHex(byte(remainder))) + hex

		n = n / 16
	}
	return hex
}

func byToHEX(b byte) string {
	result := []byte{getHex((b >> 4) & 0x0F), getHex((b & 0x0F))}

	return string(result)
}
