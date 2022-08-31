package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {

	var isThere bool
	var result []int
	for i, value1 := range angka {
		for j, value2 := range angka {
			if i == j {
				continue
			} else if value1 == value2 {
				isThere = false
				break
			} else {
				isThere = true
			}
		}
		if isThere {
			value, _ := strconv.Atoi(string(value1))
			result = append(result, value)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
