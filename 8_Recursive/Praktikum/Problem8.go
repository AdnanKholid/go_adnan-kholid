package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(item []string) []pair {
	temp := make(map[string]int)
	keys := make([]string, 0, len(item))

	for _, val := range item {
		temp[val]++
	}

	for key := range temp {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return temp[keys[i]] < temp[keys[j]]
	})

	var res []pair
	for _, k := range keys {
		res = append(res, pair{k, temp[k]})

	}
	return res
}
func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))

	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))

	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
