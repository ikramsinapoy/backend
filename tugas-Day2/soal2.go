package main

import "fmt"

func main() {
	var name string
	var nilai int
	var index string

	fmt.Print("Masukkan nama siswa : ")
	fmt.Scanf("%s\n", &name)
	fmt.Print("Masukkan nilai siswa : ")
	fmt.Scanf("%d\n", &nilai)
	fmt.Println("")

	switch {
	case nilai > 100 || nilai < 0:
		index = "Nilai Invalid"
	case nilai > 80:
		index = "A"
	case nilai > 65:
		index = "B"
	case nilai > 50:
		index = "C"
	case nilai > 35:
		index = "D"
	default:
		index = "E"
	}

	fmt.Printf("Nama Siswa : %s \nNilai: %s \n", name, index)
}
