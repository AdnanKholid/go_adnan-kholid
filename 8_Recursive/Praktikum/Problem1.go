package main

import (
	"fmt"
	"math"
)

func checkPrime(number int) bool {
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrtNumber; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

var prime []int

func primeX(number int) int {
	var last int
	if prime == nil {
		prime = append(prime, 2)
	}
	last = prime[len(prime)-1]

	jumlahElemen := len(prime)

	for jumlahElemen < number {
		if last%2 == 1 {
			last += 2
		} else {
			last += 1
		}
		if checkPrime(last) {
			prime = append(prime, last)
		}
		jumlahElemen = len(prime)
	}
	return prime[number-1]
}

func main() {
	fmt.Println(primeX(1))

	fmt.Println(primeX(5))

	fmt.Println(primeX(8))

	fmt.Println(primeX(9))

	fmt.Println(primeX(10))
}
