/*
 Untuk memberitahu bahwa aplikasi
 kita sedang berjalan di package main.
*/
package main

/*
	untuk mengimport package-package Go,
	baik built-in package atau pun third-party
	package. Package fmt merupakan package yang
	memiliki fungsi untuk menangani masalah
	input-output, seperti menampilkan
	sesuatu kelayar.
*/
import (
	"fmt"
)

// Fungsi utama dalam Go, fungsi ini akan otomatis di jalankan
func main() {
	// untuk menampilkan sesuatu ke layar
	fmt.Println("Hello Golang")
}
