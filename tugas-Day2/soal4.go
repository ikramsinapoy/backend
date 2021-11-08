package main

import "fmt"

func primeNumber(number int) bool {
	var count int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count += 1
		}
	}

	if count == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(primeNumber(1500450271))
	fmt.Println(primeNumber(1000000007))
	// fmt.Println(primeNumber(17))
	// fmt.Println(primeNumber(20))
	// fmt.Println(primeNumber(35))
}
