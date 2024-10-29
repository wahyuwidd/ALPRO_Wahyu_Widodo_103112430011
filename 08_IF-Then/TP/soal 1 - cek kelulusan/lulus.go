package main

import "fmt"

func main() {
	var nilai int
	var pesan = "Lulus" //inisialisasi pesan dengan nilai lulus
	fmt.Print("Nilai ujian: ")
	fmt.Scan(&nilai)
	if nilai < 70 { //jika nilai yang di input kurang dari 70 maka ubah pesan yang tadinya lulus jadi tidak lulus
		pesan = "Tidak Lulus"
	}
	fmt.Print(pesan)
}
