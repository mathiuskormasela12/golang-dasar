package main

import (
	"fmt"
)

func main() {

	/*
		============ Belajar Looping ========

		Looping adalah sebuah statement yang
		berfungsi untuk kita mengulang sesuatu
		secara terus menerus berdasarkan kondisi
		tertentu.
	*/

	// Di dalam Go kita mengenal For loop

	// 1. Fol Loop standar
	var index int = 1

	for index <= 5 {
		fmt.Println("Indexnya", index)
		index++
	}

	fmt.Println("==============")

	// 2. Short-hand
	for index := 1; index <= 5; index++ {
		fmt.Println("Indexnya", index)
	}

}
