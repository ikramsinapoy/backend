package main

import "fmt"

func BinarySearch(array []int, x int) {

	result := 0
	low := 1
	high := len(array) - 1

	for low <= high && result == 0 {
		mid := (high + low) / 2
		if x < array[mid] {
			high = mid - 1
		} else if x > array[mid] {
			low = mid + 1
		} else {
			result = mid
		}
	}
	if result == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(result)
	}
}

func main() {

	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3) // 2

	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5) // 3

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53) // 6

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1

}
