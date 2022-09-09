package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	var count = 0
	jml := 0
	for a, v := range s.score {
		count += v
		jml = a + 1
	}
	res := count / jml

	return float64(res)
}

func (s Student) Min() (min int, name string) {
	value := s.score

	for x, v := range value {
		if x == 0 || v < min {
			min = v
			name = s.name[x]
		}
	}

	return min, name
}

func (s Student) Max() (max int, name string) {
	value := s.score

	for x, v := range value {
		if v > max {
			max = v
			name = s.name[x]
		}
	}
	return max, name
}

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input " + string(i) + " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Printf("\n\nAvarage Score Students is %.2f\n", a.Avarage())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ") ")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ") ")
}
