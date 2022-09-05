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

func primaSegiEmpat(wide, high, start int) {
	lebarSegi := 1
	jumlahBilangan := 0
	jumlahPrima := 0

	for jumlahBilangan < (high * wide) {
		if start%2 == 1 && start > 2 {
			start += 2
		} else {
			start++
		}
		if checkPrime(start) {
			if lebarSegi <= wide {
				fmt.Printf("%d ", start)
				lebarSegi++
			} else {
				fmt.Println()
				fmt.Printf("%d ", start)
				lebarSegi = 2
			}
			jumlahBilangan++
			jumlahPrima += start
		}
	}
	fmt.Println()
	fmt.Printf("%d", jumlahPrima)
}
func main() {
	primaSegiEmpat(2, 3, 13)

	fmt.Printf("\n\n")

	primaSegiEmpat(5, 2, 1)
}
