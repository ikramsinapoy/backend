package main

import "fmt"

func pow(x, n int) int {

	result := 1
	if n == 0 {
		return 1
	} else {
		for i := 0; i < n; i++ {
			result = result * x
		}
	}

	return result

}

func main() {

	fmt.Println(pow(2, 3)) // 8

	fmt.Println(pow(7, 2)) // 49

	fmt.Println(pow(10, 5)) // 100000

	fmt.Println(pow(17, 6)) // 24137569

	fmt.Println(pow(5, 3)) // 125

}
