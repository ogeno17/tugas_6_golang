package main

import (
	"math/rand"
	l "niomic/golang/tugas_6_golang/library"
)

func main() {
	var n1, n2, n3 l.Mahasiswa
	n1.Nama = "Ahmad"
	n2.Nama = "Abi"
	n3.Nama = "Zaka"

	n1.Umur = rand.Intn(100)
	n2.Umur = rand.Intn(100)
	n3.Umur = rand.Intn(100)

	l.TampilMahasiswa(n1)
	l.TampilMahasiswa(n2)
	l.TampilMahasiswa(n3)

}
