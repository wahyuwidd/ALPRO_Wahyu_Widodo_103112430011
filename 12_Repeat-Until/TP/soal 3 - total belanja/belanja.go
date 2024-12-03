package main

import "fmt"

func main() {
	var harga, total int
	var selesai bool
	for selesai = false; !selesai; { //bentuk dari perulangan repeat-until
		//variabel selesai dideklarasi nilai boolean nya false agar perulangan berhenti ketika variabel selesai bernilai true
		fmt.Print("Masukkan harga barang (ketik 0 untuk selesai): ")
		fmt.Scan(&harga)
		//menjumlahkan harga yang diinputkan user
		total = total + harga
		//variabel selesai menjadi true jika harga yang diinputkan user adalah 0
		selesai = harga == 0
	}
	//ketika perulangan berhenti maka tampilkan total belanja
	fmt.Println("Total belanja Anda: ", total)
}
