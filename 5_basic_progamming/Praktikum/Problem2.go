package main

import (
	"fmt"
)

func main() {
	//deklaration
	var studentScore int
	var nama, grade string

	//inputs
	fmt.Print("nilai: ")
	fmt.Scanf("%d\n",&studentScore)
	fmt.Print("nama siswa: ")
	fmt.Scanf("s\n", &nama)

	//process: menentukan grade
	switch  {
	case studentScore >= 80 && studentScore <= 100:
		grade = "Nilai A"
	case studentScore >= 65 && studentScore <= 79:
		grade = "Nilai B"
	case studentScore >= 50 && studentScore <= 64:
		grade = "Nilai C"
	case studentScore >= 35 && studentScore <= 49:
		grade = "Nilai D"
	case studentScore >= 0 && studentScore <= 34:
		grade = "Nilai E"
	default:
		grade = "Nilai Invalid"		
	}

	fmt.Printf("Siswa %s %s\n", nama, grade)
}