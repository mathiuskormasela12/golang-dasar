package main

import (
	"fmt"
)

/*
	A. Struct

	Struct merupakan template untuk membuat
	sebuah object dalam Golang oleh sebab
	itu kita dapat menyebutkan bahwa
	Struct merupakan class di dalam Golang.

	Namun Struct hanya dapat menampung
	property-property object. Untuk method-nya
	kita definisikan di luar Struct. Selain
	ini di dalam struct kita tidak
	bisa langsung mendefinisikan nilai
	dari setiap propertynya, nilai
	setiap property akan di definisikan ketika
	menginstansiasi structnya.

	Struct juga bersifat "Value Type", maksudnya
	jika kita mengoper sebuat object yg variable lain
	dan mengubah propertynya melalui variable yang
	baru, maka property di variable yang lama
	tidak akan ikut ke ubah.

	rumus membuat struct :

	type NamaStruct struct {
		namaProperty1 tipeDataProperty1
		...dst
	}

	atau short-handnya jika ada beberapa
	property yg memiliki tipe data yang sama

	type NamaStruct struct {
		namaProperty1, namaProperty2 tipeDataProperty1dan2
	}

	rumus instansiasi struct
	var namaObject = NamaStruct{ namaProperty: nilaiProperty }
	atau
	var namaObject = NamaStruct{ valueProperty1, valueProperty2 } -> urutan propertynya harus tepat dengan urutan yg di buat di structnya

	instansiasi struct tanpa mendefinisikan value dari setiap propertynya
	var namaObject NamaStruct

	B. Interface

	Dalam Go terdapat juga interface, namun di dalam Go
	interface hanya bisa di isi oleh method-method.

	rumus :

	type namaInterface interface  {
		namaMethod() tipeData
	}
*/

// Membuat Class/Struct Mahasiswa
type Mahasiswa struct {
	nama string
	usia int
}

// Membuat interface Buku
type IBuku interface {
	tampilanHargaBuku() int
	cetak()
}

// Membuat Class/Struct BukuCoding
type BukuCoding struct {
	judul string
	harga int
}

// Membuat method-method BukuCoding
func (bukuKoding BukuCoding) cetak() {
	fmt.Println("Judul bukunya adalah", bukuKoding.judul)
}

func (bukuKoding BukuCoding) tampilanHargaBuku() int {
	return bukuKoding.harga
}

func main() {
	fmt.Println("=========== Belajar OOP di Go Language ===========")

	fmt.Println("")
	fmt.Println("1. Struct dan Method")

	// menginstansiasi struct tanpa mendefinisikan value tiap propertynya
	var mathius Mahasiswa
	fmt.Println(mathius)

	// menginstansiai struct sekalin mendefinisikan value tiap propertynya
	var mathius2 = Mahasiswa{
		nama: "Mathius",
		usia: 20,
	}

	// mencetak semua objectnya
	fmt.Println(mathius2)

	// menampilan nilai property object
	fmt.Println(mathius2.nama)

	// memanggil method setNama untuk mengubah nama
	mathius2.setNama("Mathius Kormasela")
	fmt.Println("Nama saya sekarang adalah", mathius2.nama)

	fmt.Println("")
	fmt.Println("2. Interface")

	var golangDasar = BukuCoding{judul: "Belajar Go Language Dari 0", harga: 35000}
	golangDasar.cetak()
	fmt.Printf("Harga buku '%s' adalah Rp. %d", golangDasar.judul, golangDasar.harga)
}

// Membuat Method
/*
	Ketika kita ingin mengubah value
	dari suatu property struct didalam
	method kita harus menggunakan pointer.
	Oleh karena itu kita menulis
	mahasiswa *Mahasiswa.

	Rumus membuat method :
	a. Method biasa (tidak bisa return)
	func (namaObjectnya namaStruct) namaMethod() {
		statement
	}

	b. Method biasa (bisa return)
	func (namaObject namaStruct) namaMethod() dataType {
		return
	}

	c. Method Yang bisa mengubah value property struct (tidak bisa return)
	func (namaObjectnya *namaStruct) namaMethod() {
		statement
	}

	d. Method Yang bisa mengubah value property struct (bisa return)
	func (namaObject *namaStruct) namaMethod() dataType {
		return
	}
*/
func (mahasiswa *Mahasiswa) setNama(nama string) {
	mahasiswa.nama = nama
}
