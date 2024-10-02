package main

import "fmt"

func main() {
	var r, hasil float32
	const pi = 3.14 //nilai pi lingkaran
	fmt.Scan(&r)
	hasil = pi * r * r //rumus menghitung luas lingkaran
	fmt.Println(hasil)
}
