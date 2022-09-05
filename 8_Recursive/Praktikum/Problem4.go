package main

import "fmt"

func MaxSequence(arr []int) int {
	tempTotal := 0
	totalTertinggi := 0
	for i, v := range arr {

		if (v > 0) && (tempTotal == 0) {
			tempTotal += v

		} else if (tempTotal > 0) || (i > 0) {
			tempTotal += v
		}

		if tempTotal > totalTertinggi {
			totalTertinggi = tempTotal
		}

		if tempTotal < 0 {
			tempTotal = 0
		}

	}

	return totalTertinggi
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))

	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))

	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
