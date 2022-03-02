package main

import (
	"fmt"
)

func main() {
	// var nilai int = 90

	if nilai := 70; nilai >= 90 {
		fmt.Println("Luar biasa")
	} else if nilai >= 80 && nilai < 90 {
		fmt.Println("Nice joob")
	} else {
		fmt.Println("Kurang")
	}
}
