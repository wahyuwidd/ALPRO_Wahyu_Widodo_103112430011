package main

import "fmt"

func main() {
	var angka int
	var pesan = "Ganjil" //inisialisasi pesan dengan nilai ganjil
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)
	if angka%2 == 0 { //jika angka habis dibagi 2 maka ubah pesan yang tadinya ganjil jadi genap
		pesan = "Genap"
	}
	fmt.Print("Angka adalah ", pesan)
}
