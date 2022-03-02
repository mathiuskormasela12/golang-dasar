package main

import (
	"fmt"
)

func main() {
	/*
		=========== Belajar Array ===========

		Array merupakan tipe data yang
		dapat memiliki lebih dari satu
		value.

		Rumus :
		1. Cara Pertama

		Kita wajib mendeklaraiskan panjangArraynya kalo gk error

		var namaVariable [panjangArraynya]dataType
		namaVariable[0] = value1
		namaVariable[1] = value2
		...dst

		2. Cara Kedua

		Kita gk wajib mendeklarasikan panjang arraynya, bisa di kosongin atau di isi dengan ...
		jika kita belum tahu array tersebut panajngnya bakal berapa
		var namaVariable = [panjangArraynya]dataType{value1, value2}
	*/

	// 1. Mendeklaraiskan Array
	var skills [4]string
	skills[0] = "HTML"
	skills[1] = "CSS"
	skills[2] = "Javascript"
	skills[3] = "Golang"

	// 2. Mendeklarasikan sekaligus mengassign nilainya
	var frameworksSkill = []string{"Express Js", "React Js", "React Native", "Gin"}

	for index := 0; index < len(skills); index++ {
		fmt.Println(skills[index])
	}

	for index := 0; index < len(frameworksSkill); index++ {
		fmt.Println(frameworksSkill[index])
	}

	var data = [3]string{}
	// mencetak panjang array atau string
	fmt.Println(len(data))
}
