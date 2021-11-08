package main

import "fmt"

func SimpleEquations(a, b, c int) {

	x, y, z := 0, 0, 0
	for i := 1; i < c; i++ {
		for j := 1; j < c; j++ {
			for k := 1; k < c; k++ {
				if i+j+k == a && i*j*k == b && ((i*i)+(j*j)+(k*k)) == c {
					x, y, z = k, j, i
					break
				}
			}
		}
	}
	if x != 0 && y != 0 && z != 0 {
		fmt.Println(x, y, z)
	} else {
		fmt.Println("No Solution")
	}

}

func main() {

	SimpleEquations(1, 2, 3) // no solution

	SimpleEquations(6, 6, 14) // 1 2 3

}
