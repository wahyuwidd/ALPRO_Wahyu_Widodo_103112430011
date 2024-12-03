package main

import "fmt"

func main() {
	var username, password string
	var percobaan int
	fmt.Print("Masukkan username dan password: ")
	fmt.Scan(&username, &password)
	for username != "Admin" || password != "Admin" {
		fmt.Scan(&username, &password)
		percobaan++
	}
	fmt.Println(percobaan, "percobaan gagal login")
}
