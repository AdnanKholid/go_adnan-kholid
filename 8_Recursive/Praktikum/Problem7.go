package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	var result interface{} = "tutup kartu"

	var deckMax int
	if deck[0] > deck[1] {
		deckMax = deck[0]
	} else {
		deckMax = deck[1]
	}

	for i := 0; i < len(cards); i++ {
		cardsArr := cards[i]

		if cardsArr[0] == deckMax || cardsArr[1] == deckMax {
			result = cardsArr
			return result
		}
		if i+1 >= len(cards) {
			if deckMax == deck[0] {
				deckMax = deck[1]
			} else {
				deckMax = deck[0]
			}
		}
	}
	return result
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))

	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}
