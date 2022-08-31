package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {

	mun := make(map[string]int)
	for _, val := range angka {
		val := string(val)
		if val != " " {
			mun[val]++
		}
	}

	var output []int
	for q, val := range mun {
		if val == 1 {
			check, _ := strconv.Atoi(q)
			output = append(output, check)

		}
	}

	fmt.Println("Input ", angka)
	fmt.Println("\n Output: ")

	return output
}

func main() {
	fmt.Println(munculSekali("1234123"))

	fmt.Println(munculSekali("76523752"))

	fmt.Println(munculSekali("12345"))

	fmt.Println(munculSekali("1122334455"))

	fmt.Println(munculSekali("0872504"))
}
