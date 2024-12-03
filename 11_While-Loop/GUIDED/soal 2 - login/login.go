package main

import "fmt"

func main() {
	var input, token string
	fmt.Print("Masukkan password: ")
	fmt.Scan(&input)
	token = "12345abcde"
	for input != token {
		fmt.Scan(&input)
	}
	fmt.Print("Selamat Anda berhasil login")
}
