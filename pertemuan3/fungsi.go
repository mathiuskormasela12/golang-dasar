package main

import (
	"fmt"
)

/*
	============== Belajar Function ==============

	A. Pengertian

	   Function merupakan kumpulan kode-kode
		 yang dapat kita panggil berulang-ulang.

		 Jenis - Jenis Function :

		 1. Function with single return

		 func namaFungsi() dataType {
			 // statement
			 return
		 }

		 2. Function with multiple return

		 func namaFungsi() (dataType, dataType) {
			 // statement
			 return value1, value2
		 }

		 3. Function with multiple return shorthand

		 func namaFungsi() (namaVariable dataType, namaVariable dataType) {
			 // statement
			 return
		 }

		 4. Function without return

		 func namaFungsi() {
			 // statement
		 }
*/

func fungsiTanpaReturn(data string) {
	fmt.Println("Ini " + data + " Tanpa Return")
}

func fungsiDenganReturn() string {
	return "Ini Fungsi dengan return"
}

func fungsiDenganMultipleReturn() (string, int, bool) {
	return "Mathius", 20, false
}

func fungsiShortHandReturn() (nama string, usia int) {
	// assign nilai ke  variable bukan inisialisasi
	nama = "Mathius"
	usia = 20

	return
}

func main() {
	fmt.Println("Hello Go")
	fungsiTanpaReturn("Fungsi")
	fmt.Println(fungsiDenganReturn())

	var nama, usia, menikah = fungsiDenganMultipleReturn()
	fmt.Println("Nama saya", nama)
	fmt.Println("usia saya", usia)

	if menikah {
		fmt.Println("Sudah menikah")
	} else {
		fmt.Println("Belum menikah")
	}

	fmt.Println("==================================")
	var nama2, usia2 = fungsiShortHandReturn()
	fmt.Println(nama2, usia2)
}
