package main

import (
	"fmt"
	"strconv"
)

/*
   =============== Pertemuan 2 =========

	 A. Belajar Variable

	    Variable merupakan sebuah tempat untuk
			menyimpan sebuah nilai atau cara kita
			untuk memberi nama kepada sebuah nilai
			sehingga ketika kita ingin memanggil
			nilai tersebut kita hanya perlu memanggil
			nama variablenya.

			rumus :
			var namaVariable = nilainya
			atau
			var namaVaribale tipeData = nilainya

	 B. Menulis Variable kosong dan Menyingkat Penulisan Beberapa Variable

	    1. Menulis Variable kosong
	    Defaultnya dalam Go, ketika kita ingin
			membuat sebuah variable, kita wajib
			untuk memberi nilai kepada variable tersebut
			kalau tidak maka akan error.

			rumus :
			var namaVariable = ""
			var namaVariable tipeData

			2. Menyingkat Penulisan Beberapa Variable
			Ketika kita ingin membuat beberapa variable
			kita bisa menyingkatnya.

			rumus :
			var variablePertama, variableKedua = nilaiVariablePertama, nilaiVariableKedua
			var variablePertama, variableKedua tipeData = nilaiVariablePertama, nilaiVariableKedua -> jadi semua variablePertama dan variableKedua tipe datanya harus sama tidak boleh beda, kalau tidak maka akan error

	 C. Menulis Commentar

	   Commentar berfungsi untuk memberi keterangan
		 mengenai suatu kode, seperti memberitahu siapa
		 yang menulis, siapa yang mengedit dan logic atau
		 tujuan dari kode nya itu apa. Commentar juga
		 berfungsi untuk mematikan sebuah kode. Commentar
		 tidak akan di jalankan atau di compile oleh Golang.
		 Jadi commentar benar-benar hanya untuk di baca
		 oleh si programmer atau pun siapapun yg sedang
		 membaca codenya.

	 D. Penulisan Singkat Variable

	    Kita dapat menulis sebuah variable dengan
			lebih singkat di Go, maksudnya tanpa keyword
			var. Namun ini hanya bisa di lakukan di
			dalam sebuah function, jika kita menulis
			diluar function maka kode kita akan error.
			Namun kita tidak bisa mendeklarasikan tipe datanya
			secara explisit.

			rumus :
			namaVariable := nilainya

		E. Membuat Konstanta Variable

		   Untuk membuat constanta variable di dalam
			 Go kita bisa menggunakan keyword const.
			 Constanta variable adalah sebuah variable
			 yang nilainya tidak bisa di ubah.

			 rumus :
			 const namaVariable = nilai


		Notes :
		1. Go tidak bisa re-assign nilai dari sebuah variable dengan tipe data yg berbeda, karena Go itu strict
		2. Jika kita mendeklarasikan sebuah variable, maka kita harus menggunakan variable tersebut, jika tidak maka akan error.


*/

// Jenis-Jenis Commentar di Golang
// Single Line Commentar
/*
 Multi Line
 Commentar
*/

// var nama = "Mathius"
// var marga string = "Kormasela"

var nama, marga = "Mathius", "Kormasela"

// error
// language := "Go"

func main() {
	// tipe data string
	language := "GO"

	// tipa data bilangan bulat bulat
	var age int = 20

	// tipe data bilangan desimal
	var val float32 = 20.9
	var val2 float64 = 300.2

	// tipe data boolean
	var married bool = false

	// Operasi Aritmatika
	var x, y, z int = 10, 20, 3
	var perkalian = x * y
	var pembagian = x / y
	var penjumlahan = x + y
	var pengurangan = x - y
	var modulus = x % z

	// increment menambah satu
	var k = 10
	k++

	// decrement menguarangi satu
	var j = 5
	j--

	const gaji = 6500000
	// error
	// gaji = 20

	fmt.Println("Hello", nama, marga)
	fmt.Println("I'm learning " + language)
	fmt.Println("Saya berusia", age, "tahun")
	fmt.Println(val)
	fmt.Println(val2)
	fmt.Println(married)
	fmt.Println(pengurangan)
	fmt.Println(penjumlahan)
	fmt.Println(pembagian)
	fmt.Println(perkalian)
	fmt.Println(modulus)
	fmt.Println(k)
	fmt.Println(j)
	fmt.Println("Gaji saya Rp.", gaji)

	var i int = 10
	var name string = "20"
	var nameInt, _ = strconv.Atoi(name)

	/*
		  var nameInt, _ = strconv.Atoi(name)

			keteranga :
			_ = ini adalah variable untuk menampung error,
			    jika ada error ketika sedang melakuian konversi.
					Kita menulis nama variable nya deng tanda "_" agar
					tidak terjadi error walaupun variable "_" tidak
					kita tampilakan, klo kita tulis dengan nama lain,
					otomatis kita harus selalu mencetak errornya kalo gk
					bakal error
	*/

	// Konversi tipe data integer ke string
	fmt.Println("Nilai saya " + strconv.Itoa(i))
	// Konversi tipe data string ke integer
	fmt.Println(nameInt * 20)
}
