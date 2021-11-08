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

	fmt.Println(primeNumber(1000000007)) // true

	fmt.Println(primeNumber(1500450271)) // true

	fmt.Println(primeNumber(1000000000)) // false

	fmt.Println(primeNumber(10000000019)) // true

	fmt.Println(primeNumber(10000000033)) // true

}
