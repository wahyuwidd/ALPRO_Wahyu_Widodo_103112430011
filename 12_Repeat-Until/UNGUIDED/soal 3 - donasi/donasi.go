package main

import "fmt"

func main() {
	var target, donasi, total, jumlah_donatur int
	var selesai bool
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)
	for selesai = false; !selesai; {
		fmt.Print("Masukkan jumlah donasi: ")
		fmt.Scan(&donasi)
		total += donasi
		selesai = total >= target
		jumlah_donatur++
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlah_donatur, donasi, total)
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.", total, jumlah_donatur)

}
