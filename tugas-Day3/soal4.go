package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	var result []int

	var check = make(map[rune]int)
	for _, val := range angka {
		check[val] = check[val] + 1
	}

	for number := range check {
		if check[number] == 1 {
			// valuenya masih berbentuk ascii oleh karena itu perlu dikurangi 48
			result = append(result, int(number)-48)
		}
	}
	return result

}

func main() {

	fmt.Println(munculSekali("1234123")) // [4]

	fmt.Println(munculSekali("76523752")) // [6, 3]

	fmt.Println(munculSekali("12345")) // [1 2 3 4 5]

	fmt.Println(munculSekali("1122334455")) // []

	fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]

}
