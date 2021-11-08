package main

import (
	"fmt"
	"time"
)

func letterFreq(input string) {
	mapping := map[rune]int{}

	go func(str map[rune]int) {
		for _, value := range input {
			mapping[value] += 1
		}
	}(mapping)

	time.Sleep(100 * time.Millisecond)

	for key, value := range mapping {
		fmt.Printf("%c => %d\n", key, value)
	}
}

func main() {
	input := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	letterFreq(input)
}
