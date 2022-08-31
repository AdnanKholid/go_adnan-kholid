package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	res := make([]string, 0)
	inRes := make(map[string]bool)
	x := append(arrayA, arrayB...)

	for _, i := range x {
		if _, ok := inRes[i]; !ok {
			inRes[i] = true
			res = append(res, i)
		}
	}

	return res
}

func main() {

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"segel", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))
}
