package main

import "fmt"

func MaximumBuyProduct(money int, productPrice []int) {
	temp := make([]int, 0)
	res := 0
	for i := 0; i < len(productPrice)-1; i++ {
		for j := 0; j < len(productPrice)-i-1; j++ {
			if productPrice[j] > productPrice[j+1] {
				productPrice[j], productPrice[j+1] = productPrice[j+1], productPrice[j]
			}
		}
	}

	for i := 0; i < len(productPrice); i++ {
		if money >= productPrice[i] {
			money -= productPrice[i]
			temp = append(temp, i)
		}
	}
	for j := 1; j <= len(temp); j++ {
		res = j
	}
	fmt.Println(res)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})

	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})

	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})

	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})

	MaximumBuyProduct(0, []int{10000, 30000})
}
