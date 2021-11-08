package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	merge := []string{}
	check := make(map[string]bool)
	result := []string{}

	merge = append(arrayA, arrayB...)

	for _, data := range merge {
		if _, value := check[data]; !value {
			check[data] = true
			result = append(result, data)
		}
	}
	
	return result
}

func main() {
	
	// Test cases

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))

	// []

}
