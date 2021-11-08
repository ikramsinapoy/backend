package main

import (
	"fmt"
	"sort"
)

func minimizeBias(ratings []int32) int32 {
	number := int(ratings)
	sort.Ints(number)

	return int32(number)
}

func main() {
	ratings := []int32{3, 5, 3, 1, 3, 5, 62, 1, 52, 5, 1}
	fmt.Println(minimizeBias(ratings))
}
