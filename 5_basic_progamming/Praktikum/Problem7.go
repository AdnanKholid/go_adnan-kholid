package main

import "fmt"

func main() {
	var row int = 5
	var k int

	for i := 1; i <= row; i++ {
		k = 0
		for space := 1; space <= row-i; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}
}
