package main

import "fmt"

func main() {
	var f, hasil int
	fmt.Scan(&f)
	hasil = (f - 32) * 5 / 9 //rumus konversi suhu fahrenheit ke celsius
	fmt.Println(hasil)
}
