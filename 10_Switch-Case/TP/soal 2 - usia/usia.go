package main

import "fmt"

func main() {
	var usia int
	// Meminta input usia dari pengguna
	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)

	// Menentukan kategori usia menggunakan switch
	switch {
	case usia >= 0 && usia <= 12:
		//jika usia yg di input user lebih dari 0 dan kurang dari sama dengan 12 tampilkan outpout kategori anak anak
		fmt.Println("Kategori: Anak-anak")
	case usia >= 13 && usia <= 17:
		//jika usia yg di input user >= 13 dan <= 17 tampilkan output kategori remaja
		fmt.Println("Kategori: Remaja")
	case usia >= 18 && usia <= 64:
		//jika usia yg di input user >= 18 dan <= 64 tampilkan output kategori dewasa
		fmt.Println("Kategori: Dewasa")
	case usia >= 65:
		//jika usia yg di input user >= 65 tampilkan output kategori lansia
		fmt.Println("Kategori: Lansia")
	default:
		//jika tidak ada kondisi di atas yg terpenuhi tampilkan output usia tidak valid.
		fmt.Println("Usia tidak valid.")
	}
}
