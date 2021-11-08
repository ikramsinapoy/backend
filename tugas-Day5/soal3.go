package main

import "fmt"

func primeX(number int) bool {

	check := 0
	if number < 2 {
		return false
	} else if number == 2 {
		return true
	} else {
		for i := 1; i <= number; i++ {
			if number%i == 0 {
				check += 1
			}
		}
		if check == 2 {
			return true
		} else {
			return false
		}
	}
}

func primaSegiEmpat(high, wide, start int) {

	jumlahPrime := high * wide
	arrPrime := []int{}
	total := 0

	for i := start + 1; len(arrPrime) < jumlahPrime; i++ {
		if primeX(i) == true {
			arrPrime = append(arrPrime, i)
		}
	}

	j := 0
	for j < len(arrPrime) {
		for h := 0; h < wide; h++ {
			for w := 0; w < high; w++ {
				fmt.Print(arrPrime[j], " ")
				total = total + arrPrime[j]
				j++
			}
			fmt.Println("")
		}
	}

	fmt.Println(total)
	fmt.Println("")
}

func main() {

	primaSegiEmpat(2, 3, 13)

	/*

	   17 19

	   23 29

	   31 37

	   156

	*/

	primaSegiEmpat(5, 2, 1)

	/*

	   2  3  5  7 11

	   13 17 19 23 29

	   129

	*/

}
