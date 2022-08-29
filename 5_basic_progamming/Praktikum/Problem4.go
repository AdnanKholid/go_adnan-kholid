package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) {
	if number < 2 {
		fmt.Println(false)
		return
	}
	prime := int(math.Sqrt(float64(number)))
	for i := 2; i <= prime; i++ {
		if number%i == 0 {
			fmt.Println(false)
			return
		}
	}
	fmt.Println(true)

}

func main() {
	primeNumber(0)
	primeNumber(13)
	primeNumber(17)
	primeNumber(20)
	primeNumber(35)
}
