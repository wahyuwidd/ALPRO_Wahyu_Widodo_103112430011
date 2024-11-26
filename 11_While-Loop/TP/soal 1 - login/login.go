package main

import "fmt"

func main() {
	var passwordLogin, password string
	var percobaan, percobaanMax int
	fmt.Print("Masukkan Password: ")
	percobaanMax = 3           //maksimal percobaan yg salah
	passwordLogin = "admin123" //password yg benar
	fmt.Scan(&password)
	//kenapa ada -1? karena untuk mengurangi jumlah percobaan salah yang di awal
	for password != passwordLogin && percobaan < percobaanMax-1 { //jika password tidak sama dengan passwordLogin dan percobaan kurang dari percobaanMax maka lakukan perulangan input password
		fmt.Println("Password salah. Coba lagi")
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&password)
		//setiap percobaan salah maka percobaan akan bertambah
		percobaan++
	}
	if password == passwordLogin {
		//jika password yg dimasukkan benar tampilkan login berhasil karena password yg di masukkan benar
		fmt.Println("Login berhasil")
	} else {
		//jika tidak maka tampilkan "Login ditolak"
		fmt.Println("Login ditolak")
	}
}
