package main

import "fmt"

func main() {
	var input, point, pointTambah, Jumlah int
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&input)
	// inisialisasi nilai awal untuk poin 10 dan poin tambahan 5
	point = 10
	pointTambah = 5
	for i := 1; i <= input; i++ {
		// iterasi untuk menghitung poin, jadi setiap iterasi nanti di kali n
		Jumlah = i * point
		// jika input lebih dari 5, poin tambahan akan di tambah menajadi n
		if input > 5 {
			Jumlah = (5 * point) + ((input - 5) * pointTambah * 3)
		}
	}
	fmt.Print("Total poin yang didapat: ", Jumlah, " poin")

}
