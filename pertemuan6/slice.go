package main

import (
	"fmt"
)

func main() {
	/*
		============ Belajar Slice ========

		A. Pengertian

		   Slice merupkan tipe data yang mirip
			 dengan array, namun lebih fleksibel
			 di banding array. Array di Golang kita
			 wajib untuk menentukan panjangnya atau
			 kita harus menulis dengan tand ...
			 jika kita belum tahu panjangnya.
			 Kalau kita tulis kosong, maka array akan
			 error, kalau slice kita tidak perlu mendefinisikan
			 berapa panjang arraynya. Tidak perlu menulis ...
			 jika kita blm tahu panjang array nya.

			 rumus :
			 1. Cara Pertama
			 var namaVariable = []dataType{value1, value2 ..dst}

			 2. Cara Kedua
			 var namaVariable = make([]dataType, panjangnya, kapasitas?) -> kapasitas nya optional
	*/
	var myLanguages = [5]string{"Indonesia", "English", "Japanese", "Germany", "Korean"}

	// Menyalin nilai array ke dalam sebuah slice
	// rumus var namaVariable = namaVariableSlice[index awal:indexakhir] -> index awal defaultnya 0, ini mirip seperti method slice di Javascript
	// menyalin semua data arraynya -> : ini menyalin semua data arraynya
	// ketika menyalain array/atau slice dengan slice kita sebenernya tidak benar2 menyalin hanya mengarahkan nya ke array/slice utama, oleh karena itu perubahan yg kita lakukan di array/slice kedua akan berdampak di array/slice utama
	var myLanguagesSlice = myLanguages[:]

	// mengubah nilai array/slice
	myLanguagesSlice[1] = "British"

	fmt.Println("This is an array", myLanguages)
	fmt.Println("This is a slice", myLanguagesSlice)

	// untuk mencetak kapsitas dari sebuah array atau slice
	// kapasitas di hitung dari element array/slice pertama setelah di salin dengan slice
	fmt.Println(cap(myLanguagesSlice))

	fmt.Println("======== Slice with Make ======")
	newLang := make([]string, 5, 5)

	// menyalin slice (hanya bisa digunain kalau kita bikin slice pakai metode make)
	// rumus copy(mau di salin ke variable mana, slicenya dari variable mana)
	copy(newLang, myLanguages[:])
	newLang[4] = "Hangul"

	// menambah element slice -> mirip seperti method push di javascript
	// rumus var namaVariable = append(sumber slicenya, elementSliceBaru1, elementSliceBaru2, ..dst)
	newLang = append(newLang, "Ukraine")
	fmt.Println("This is an array", myLanguages)
	fmt.Println("Tjis is a array", newLang)
}
