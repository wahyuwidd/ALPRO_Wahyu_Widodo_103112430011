package main

import "fmt"

func main() {
	var kata string
	var ulang, jumlah int
	fmt.Print("Masukkan kata dan jumlah pengulangan: ")
	fmt.Scan(&kata, &ulang)
	for selesai := false; !selesai; {
		fmt.Println(kata)
		jumlah++
		selesai = jumlah >= ulang
	}
}
