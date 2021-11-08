package main

import (
	"errors"
	"fmt"
)

type student struct {
	name string

	nameEncode string

	score int
}

type Chiper interface {
	Encode() (string, error)

	Decode() (string, error)
}

func (s *student) Encode() (string, error) {

	var nameEncode = ""
	runes := []rune(s.name)
	var ascii rune
	for i := 0; i < len(s.name); i++ {
		if s.name[i] >= 97 && s.name[i] <= 122 {
			save := 97 - int(runes[i]) + 122
			ascii = rune(save)
		} else if s.name[i] >= 65 && s.name[i] <= 90 {
			save := 65 - int(runes[i]) + 90
			ascii = rune(save)
		} else {
			return "", errors.New("Input must be character 'a' to 'z' or 'A' to 'Z'")
		}
		nameEncode = nameEncode + string(ascii)
	}
	return nameEncode, nil

}

func (s *student) Decode() (string, error) {

	var nameDecode = ""
	runes := []rune(s.nameEncode)
	var ascii rune
	for i := 0; i < len(s.nameEncode); i++ {
		if s.nameEncode[i] >= 97 && s.nameEncode[i] <= 122 {
			save := 97 - int(runes[i]) + 122
			ascii = rune(save)
		} else if s.nameEncode[i] >= 65 && s.nameEncode[i] <= 90 {
			save := 65 - int(runes[i]) + 90
			ascii = rune(save)
		} else {
			return "", errors.New("Input must be character 'a' to 'z' or 'A' to 'Z'")
		}
		nameDecode = nameDecode + string(ascii)
	}
	return nameDecode, nil

}

func main() {

	var menu int

	var a = student{}

	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")

	fmt.Scan(&menu)

	if menu == 1 {

		fmt.Print("\nInput Student’s Name : ")

		fmt.Scan(&a.name)

		if val, err := c.Encode(); val != "" {
			fmt.Print("\nEncode of Student's Name " + a.name + " is : " + val)
		} else {
			fmt.Println(err.Error())
		}
	} else if menu == 2 {

		fmt.Print("\nInput Student’s Encode Name : ")

		fmt.Scan(&a.nameEncode)

		if val, err := c.Decode(); val != "" {
			fmt.Print("\nDecode of Student's Name " + a.name + " is : " + val)
		} else {
			fmt.Println(err.Error())
		}
	} else {

		fmt.Println("Wrong input name menu")

	}

}
