package main

import (
	"fmt"
)

func main() {
	/*
		========== Belajar Range ==========

		Range merupakan sebuah looping
		untuk menemapilkan array. Range
		mirip seperti looping for in di
		Javascript.

		Rumus :

		for var index, element := range variableArray {
			statement
		}

		keterengan :
		index -> index tiap element array
		element -> element arraynya
	*/
	var family = [...]string{"Mathius", "Tessa"}

	for index, element := range family {
		// %d untuk mencetak index yg berupa angka
		// %s untuk mencetak element yg berupa string
		fmt.Printf("%d = %s", index, element)
	}
}
