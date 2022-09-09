package main

import "fmt"

func Compare(a, b string) string {

	if a == "" || b == "" {
		return ""
	}

	temp := ""

	if len(a) > len(b) {
		a, b = b, a
	}
	for key, _ := range a {
		if a[key] == b[key] {
			temp += string(a[key])
		} else {
			break
		}
	}

	return temp
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))

	fmt.Println(Compare("KANGOORO", "KANG"))

	fmt.Println(Compare("KI", "KIJANG"))

	fmt.Println(Compare("KUPU-KUPU", "KUPU"))

	fmt.Println(Compare("ILALANG", "ILA"))
}
