package main

import "fmt"

func main() {
	var nama, nim, kelas string
	fmt.Print("Masukan Nama: ")
	fmt.Scan(&nama) //fmt.Scan buat menerima input dari user
	fmt.Print("Masukan Nim: ")
	fmt.Scan(&nim)
	fmt.Print("Masukan Kelas: ")
	fmt.Scan(&kelas)
	fmt.Println("Perkenalkan saya adalah", nama, "salah satu mahasiswa Prodi S1-IF dari kelas", kelas, "dengan NIM", nim)
}
