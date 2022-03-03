package main

import (
	"fmt"

	"github.com/mathiuskormasela12/pertemuan9/data"
)

func main() {
	fmt.Println("========= Belajar Import dan Export Module =========")

	for _, v := range data.Mahasiswa {
		fmt.Println(v)
	}
}
