package main

import (
	"fmt"
	"strconv"
)

/*
	Rumus :
	1. Membuat Empty Map
	var namaVariable = make(map[tipeDataKey]tipeDataValue)
	atau
	var namaVariable = map[tipeDataKey]tipeDataValue{}

	2. Membuat Map dan mengisi valuennya langsung
	var namaVariable = map[tipeDataKey]tipeDataValue{
		namaKey: value,
	}
*/

var mhs = map[string]string{
	"nama": "Mathius",
	"usia": "19",
}

func belajarMenghapusElementMap() {
	fmt.Println("Belajar Menghapus Element Map")

	var matkul = map[int]string{
		10291: "Pemrograman Web",
		2411:  "Basis Data",
		5131:  "MTK",
	}

	fmt.Println("")
	fmt.Println("Menampilakan Semua Matkul")

	for prop, val := range matkul {
		fmt.Println(strconv.Itoa(prop) + " = " + val)
	}

	fmt.Println("")
	fmt.Println("Menampilakan Semua Matkul Tapi Satu Matkul Udah Di Hapus")

	// menghapus sebuah element map
	// rumus delete(sumber map-nya, key-nya)
	delete(matkul, 2411)

	for prop, val := range matkul {
		fmt.Println(strconv.Itoa(prop) + " = " + val)
	}
}

func main() {
	fmt.Println("======== Belajar Tipe Data Map ===========")
	fmt.Println("Tipe data map adalah tipe data key value dalam Golang, mirip seperti object di Javascript dan array assosiatif di PHP. Untuk menampilkan seluruh nilanya kita bisa pakei range loop")
	fmt.Println("Map bersifat 'Reference Type', dimana ketika kita menyalin sebuah map ke dalam variable lain dan mengubah sesuatu melalu variable lain itu, maka map atau variable utama akan ikut kerubah. ")

	// Menampilakan map dengan range
	for key, value := range mhs {
		fmt.Println(key + " => " + value)
	}

	// Mengecek apakah di suatu map ada key tertantu
	if isExists("usia") {
		fmt.Println("Ada")
	} else {
		fmt.Println("Tidak ada")
	}

	belajarMenghapusElementMap()
}

func isExists(key string) bool {
	_, exists := mhs[key]
	return exists
}
