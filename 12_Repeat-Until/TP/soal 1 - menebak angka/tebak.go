package main

import "fmt"

func main() {
	var input int
	var selesai bool
	fmt.Print("Tebak angka (1-10): ")
	fmt.Scan(&input)
	for selesai = false; !selesai; { //bentuk dari perulangan repeat-until
		//variabel selesai dideklarasi nilai boolean nya false agar perulangan berhenti ketika variabel selesai bernilai true
		fmt.Println("Tebakan anda salah, coba lagi. ")
		fmt.Print("Tebak angka (1-10): ")
		fmt.Scan(&input)
		//disini variabel selesai menjadi true jika input/tebakan dari user adalah 4
		//nah kalo belum == 4 maka perulangan akan berjalan terus dimana nanti user akan terus menginputkan tebakan angka nya sampai benar
		selesai = input == 4 //misalkan angka rahasia nya adalah 4
	}
	//ketika perulangan berhenti yang berarti tebakan nya udah benar maka print selamat, tebakan anda benar!
	fmt.Print("Selamat, tebakan anda benar!")
}
