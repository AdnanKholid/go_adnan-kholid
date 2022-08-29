package main

import (
	"fmt"
)

func main() {
	//declaration
	var bilangan int

	//input
	fmt.Scanf("%d", &bilangan)

	//process: menentukan faktor
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}
}