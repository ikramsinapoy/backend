package main

import "fmt"

var TopDown map[int]int = map[int]int{}

func fibo(n int) int {

	if value, check := TopDown[n]; check {
		return value
	}

	if n <= 1 {
		TopDown[n] = n
	} else {
		TopDown[n] = fibo(n-1) + fibo(n-2)
	}

	return TopDown[n]

}

func main() {

	fmt.Println(fibo(0)) // 0

	fmt.Println(fibo(1)) // 1

	fmt.Println(fibo(2)) // 1

	fmt.Println(fibo(3)) // 2

	fmt.Println(fibo(5)) // 5

	fmt.Println(fibo(6)) // 8

	fmt.Println(fibo(7)) // 13

	fmt.Println(fibo(9)) // 13

	fmt.Println(fibo(100)) // 55

}
