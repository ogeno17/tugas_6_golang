package library

import "fmt"

// Mahasiswa is object
type Mahasiswa struct {
	Nama string
	Umur int
}

// TampilMahasiswa is func
func TampilMahasiswa(m Mahasiswa) {
	fmt.Println("Nama: ", m.Nama, ", Umur: ", m.Umur)
}
