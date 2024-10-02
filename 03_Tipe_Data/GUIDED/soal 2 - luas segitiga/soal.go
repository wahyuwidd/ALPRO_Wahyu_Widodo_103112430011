package main

import "fmt"

func main() {
	var alas, tinggi, luas float32
	fmt.Print("Masukan alas = ")
	fmt.Scan(&alas)
	fmt.Print("Masukan tinggi = ")
	fmt.Scan(&tinggi)
	luas = 0.5 * alas * tinggi
	fmt.Println("Luas = ", luas)
}
