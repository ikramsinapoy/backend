package main

import (
	"fmt"
	"strings"
)

func getMinDeletions(s string) int32 {
	check := make(map[string]bool)
	// arr := []string{}
	result := []string{}
	arrbaru := strings.Split(s, "")

	// for _, v := range s {
	// 	arr = append(arr, string(v))
	// }
	for _, data := range arrbaru {
		if _, value := check[data]; !value {
			check[data] = true
			result = append(result, data)
		}
	}

	return int32(len(s) - len(result))

}

func main() {
	fmt.Println(getMinDeletions("abcabc"))
}
