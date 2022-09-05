package main

import "fmt"

func FindMinAndMax(arr []int) string {
	var result string
	minIndex := 0
	maxIndex := 0

	for i, v := range arr {
		if arr[minIndex] > v {
			minIndex = i
		} else if arr[maxIndex] < v {
			maxIndex = i
		}
	}
	result = fmt.Sprintf("min: %d index: %d max: %d index: %d", arr[minIndex], minIndex, arr[maxIndex], maxIndex)
	return result
}

func main() {
	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))

	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))

	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))

	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))

	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -2}))
}
