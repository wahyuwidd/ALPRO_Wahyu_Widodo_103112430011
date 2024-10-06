package main

import "fmt"

func main() {
	var r, luas, keliling float64
	var pi float64 = 3.14 //nilai pi lingkaran
	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scan(&r)

	luas = pi * r * r     //rumus luas lingkaran
	keliling = 2 * pi * r //rumus keliling lingkaran

	//menampilkan luas dan keliling lingkaran
	fmt.Println("Luas lingkaran:", luas)
	fmt.Println("Keliling lingkaran:", keliling)
}
