package main

import (
	"fmt"
	"math"
)

func pow(x, n int) int {
	a := float64(n)
	b := float64(x)

	mathpow := math.Pow(b, a)

	return int(mathpow)
}

func main() {

	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))

}
