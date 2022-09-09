package main

import "fmt"

func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a, b *int) {
	newB := *a
	*a = *b
	*b = newB
}
