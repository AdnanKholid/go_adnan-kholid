package main

import "fmt"

func main() {

	var bilangan int
	fmt.Print("input :")
	fmt.Scan(&bilangan)
	y := bilangan + 1
	for vertical := 1; vertical < y; vertical++ {

		for horizontal := 1; horizontal < y; horizontal++ {

			fmt.Printf("%d ", horizontal)

		}

		fmt.Println()
	}

}
