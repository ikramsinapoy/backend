package main

import (
	"fmt"
	"strings"
)

// func Compare(a, b string) string {

// 	lenA := len(a)
// 	lenB := len(b)

// 	if lenA > lenB {
// 		return b
// 	} else {
// 		return a
// 	}

// }

func main() {

	// fmt.Println(Compare("AKA", "AKASHI")) // AKA

	// fmt.Println(Compare("KANGOORO", "KANG")) // KANG

	// fmt.Println(Compare("KI", "KIJANG")) // KI

	// fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU

	// fmt.Println(Compare("ILALANG", "ILA")) // ILA
	fmt.Println(strings.Contains("code", "o"))
	fmt.Println(strings.Contains("A computer science portal", "science"))
}
