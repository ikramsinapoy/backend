package main

import "fmt"

func main() {
	var T float64
	var r float64
	pi := 3.14

	fmt.Print("Masukkan tinggi tabung : ")
	fmt.Scanf("%g\n", &T)
	fmt.Print("Masukkan jari-jari : ")
	fmt.Scanf("%g\n", &r)

	p := T + r
	L := (2 * pi * r * p)
	fmt.Println(L)
}
