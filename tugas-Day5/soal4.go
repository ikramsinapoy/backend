package main

import "fmt"

func MaxSequence(arr []int) int {

	total := 0
	result := -1
	for i := 0; i < len(arr); i++ {
		total = total + arr[i]
		if total > result {
			result = total
		}
		if total < 0 {
			total = 0
		}
	}

	return result

}

func main() {

	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7

	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8

	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12

}
