package main

import (
	"fmt"
	"math"
)

func main() {
	var bilangan int
	var pangkat int
	fmt.Print("bilangan : ")
	fmt.Scan(&bilangan)
	fmt.Print("pangkat : ")
	fmt.Scan(&pangkat)

	var rumus = math.Pow(float64(bilangan), float64(pangkat))

	fmt.Printf("Hasil %d pangkat %d adalah %F", bilangan, pangkat, rumus)
}
