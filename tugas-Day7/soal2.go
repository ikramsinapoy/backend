package main

import "fmt"

func fibo(n int) int {

	fibBottomUp := make([]int, n+2)
	fibBottomUp[0] = 0
	fibBottomUp[1] = 1
	for i := 2; i < n+1; i++ {
		fibBottomUp[i] = fibBottomUp[i-1] + fibBottomUp[i-2]
	}

	return fibBottomUp[n]

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

	fmt.Println(fibo(10)) // 55

}
