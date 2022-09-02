package main

import (
	"fmt"
)

func palindrome(input string) bool {
	lengthString := len(input)
	for i := 0; i <= lengthString/2; i++ {
		if input[i] != input[(lengthString-1)] {
			return false
		}
	}

	return false

}

func main() {
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("mistar"))
	fmt.Println(palindrome("lion"))
}
