package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	//menggunakan perulangan for-loop karena jumlah iterasi/berhenti nya di ketahui
	for i := 1; i <= n; i++ { //inisialisasikan i nya 1, kemudian selama i masih kurang dari n maka lakukan increment di variabel i dan lakukan aksi nya
		//nah disini ada kombinasi IF-Then yaitu untuk menentukan bilangan prima atau bukan
		isPrime := true // Asumsi awal bahwa i adalah bilangan prima
		if i > 1 {      //karena 1 bukan bil prima
			for j := 2; j*j <= i; j++ { // Periksa faktor dari 2 hingga akar kuadrat i
				if i%j == 0 {
					isPrime = false // Jika ditemukan pembagi, maka bukan bilangan prima
					break
				}
			}
			if isPrime {
				fmt.Print(i, " ") // tampilkan jika bilangan prima
			}
		}
	}
}
