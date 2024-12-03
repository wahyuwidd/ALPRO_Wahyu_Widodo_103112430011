package main

import "fmt"

func main() {
	var kata string
	var selesai bool
	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)
	for selesai = false; !selesai; { //bentuk dari perulangan repeat-until
		//variabel selesai dideklarasi nilai boolean nya false agar perulangan berhenti ketika variabel selesai bernilai true
		fmt.Println("Anda mengetik: ", kata)
		fmt.Print("Masukkan kata: ")
		fmt.Scan(&kata)
		//variabel selesai menjadi true jika kata dari user adalah telkom
		//nah kalo belum == telkom maka perulangan akan berjalan terus dimana nanti user akan terus mengkatakan kata nya sampai kata telkom
		selesai = kata == "telkom" || kata == "TELKOM" || kata == "Telkom"
	}
	//ketika perulangan berhenti maka tampilkan program selesai
	fmt.Print("Program selesai.")
}
