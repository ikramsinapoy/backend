package main

import "fmt"

func caesar(offset int, input string) string {

	var result string

	for i := 0; i < len(input); i++ {

		offset = offset % 26
		ascii := rune(input[i]) + rune(offset)

		if ascii > 122 {
			ascii = ascii - 122
			ascii = ascii + 96
		}

		result = result + string(ascii)
	}

	return result
}

func main() {

	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // cnvc

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

}
