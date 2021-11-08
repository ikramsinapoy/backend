package main

import "fmt"

func primeX(number int) int {
	primeNumber := []int{}
	index := 0

	for i := 2; i <= 30; i++ {
		check := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				check += 1
			}
		}

		if check == 2 {
			primeNumber = append(primeNumber, i)
		}
	}

	for i, value := range primeNumber {
		if i == number-1 {
			index = value
		}
	}
	return index
}

func main() {

	fmt.Println(primeX(1)) // 2

	fmt.Println(primeX(5)) // 11

	fmt.Println(primeX(8)) // 19

	fmt.Println(primeX(9)) // 23

	fmt.Println(primeX(10)) // 29

}
