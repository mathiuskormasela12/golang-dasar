package main

import (
	"fmt"
)

// nama *string => membuat pointer baru
func coba(nama *string) {
	// mengubah nilai dari variable nama langsung dri memory
	*nama = "Matthew"
}

func main() {
	/*
		========= Belajar Pointer =======

		Untuk mengarahkan sebuah variable ke dalam variable
		yang disempen di memory.

		rumus membuat pointer :

		var namaVariable *dataType =  *namaVariableLain

		keterangan :
		* => untuk mengambil nilai dri variable nya dan untuk inisialisasi pointer
		& => untuk mengambil alamat variable nya
	*/
	var nama = "Mathius"

	// mengirim alamat variable nama
	coba(&nama)
	fmt.Println(nama)
}
