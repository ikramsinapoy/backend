package main

import "fmt"

type Student struct {
	name []string

	score []int
}

func (s Student) Avarage() float64 {
	result := 0

	for _, value := range s.score {
		result += value
	}

	return float64(result / len(s.score))
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	index := 0
	for i, value := range s.score {
		if value < min {
			min = value
			index = i
		}
	}

	return min, s.name[index]
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	index := 0
	for i, value := range s.score {
		if max < value {
			max = value
			index = i
		}
	}

	return max, s.name[index]

}

func main() {

	var a = Student{}

	for i := 0; i < 5; i++ {

		var name string

		fmt.Print("Input " + string(i) + " Student’s Name : ")

		fmt.Scan(&name)

		a.name = append(a.name, name)

		var score int

		fmt.Print("Input " + name + " Score : ")

		fmt.Scan(&score)

		a.score = append(a.score, score)

	}

	fmt.Println("\n\nAvarage Score Students is", a.Avarage())

	scoreMax, nameMax := a.Max()

	fmt.Println("Max Score Students is : "+nameMax+"(", scoreMax, ")")

	scoreMin, nameMin := a.Min()

	fmt.Println("Min Score Students is : "+nameMin+"(", scoreMin, ")")

}
