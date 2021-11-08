package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {

	result := 0
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)
	dragonDead := 0

	if len(dragonHead) < len(knightHeight) {
		for i := 0; i < len(dragonHead); i++ {
			for j := 0; j < len(knightHeight); j++ {
				if dragonHead[i] <= knightHeight[j] { // lakukan pengecekan jika sudah memotong atau belum atau bisa di delete
					dragonDead++
					result = result + knightHeight[j]
					break
				}
			}
		}
		if dragonDead == len(dragonHead) {
			fmt.Println(result)
		} else {
			fmt.Println("knight fall")
		}

	} else {
		fmt.Println("knight fall")
	}

}

func main() {

	DragonOfLoowater([]int{5, 6}, []int{7, 8, 4}) // 11

	DragonOfLoowater([]int{5, 10}, []int{5}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10

}
