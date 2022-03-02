package main

import (
	"fmt"
)

func main() {
	var nilai int = 80

	switch {
	case nilai >= 90:
		fmt.Println("Luar biasa")
	case nilai >= 80 && nilai < 90:
		fmt.Println("Keren")
	default:
		fmt.Println("Kurang")
	}

	var lang = "Python"

	switch lang {
	case "Go":
		fmt.Println("Belajar Go")
	case "Javascript":
		fmt.Println("Belajar Javascript")
	default:
		fmt.Println("Tidur")
	}
}
