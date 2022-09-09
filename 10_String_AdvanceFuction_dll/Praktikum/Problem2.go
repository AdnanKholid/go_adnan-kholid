package main

import "fmt"

func caesar(offset int, input string) string {
	var res string

	tempInput := []rune(input)

	for _, val := range tempInput {
		val = changeAscii(val, offset)
		res += string(val)
	}

	return res
}

func changeAscii(r rune, offset int) rune {
	return (((r - 'a') + rune(offset)) % 26) + 'a'
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))

}
