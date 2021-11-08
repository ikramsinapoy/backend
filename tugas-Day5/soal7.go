package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {

	total := 0
	result := make([]int, 2)
	for _, vcards := range cards {
		for _, vdeck := range deck {
			if vcards[0] == vdeck || vcards[1] == vdeck {
				if vcards[0]+vcards[1] > total {
					total = vcards[0] + vcards[1]
					result[0], result[1] = vcards[0], vcards[1]
				}
			}
		}
	}

	if total != 0 {
		return result
	} else {
		return "Tutup kartu"
	}

}

func main() {

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))

	// [3, 4]

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))

	// [6 5]

	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))

	// “tutup kartu”

}
