package main

import "fmt"

func palindrome(kata string) bool {
	for i := 0; i < len(kata)/2; i++ {
		if kata[i] != kata[len(kata)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(palindrome("civic"))       // true
	fmt.Println(palindrome("katak"))       // true
	fmt.Println(palindrome("kasur rusak")) // true
	fmt.Println(palindrome("mistar"))      // false
	fmt.Println(palindrome("lion"))        // false
}
