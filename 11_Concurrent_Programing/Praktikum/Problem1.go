package main

import "fmt"

func main() {

	temp := make(map[string]int)

	kalimat := "adnan kholid"

	go func() {
		for _, char := range kalimat {
			temp[string(char)]++
		}
	}()
	fmt.Println(temp)
}
