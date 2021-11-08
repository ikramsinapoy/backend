package main

import "fmt"

func pangkat(base, pangkat int) int {
	hasil := 1
	for i := 0; i < pangkat; i++ {
		hasil = hasil * base
	}
	return hasil

}

func main() {
	fmt.Println(pangkat(2, 3)) // 8

	fmt.Println(pangkat(5, 3)) // 125

	fmt.Println(pangkat(10, 2)) // 100

	fmt.Println(pangkat(2, 5)) // 32

	fmt.Println(pangkat(7, 3)) // 343
}
