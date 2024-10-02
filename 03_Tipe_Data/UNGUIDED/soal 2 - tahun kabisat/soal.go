package main

import "fmt"

func main() {
	var tahun int
	var isKabisat bool

	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)
	isKabisat = tahun%4 == 0
	fmt.Println("Kabisat:", isKabisat)
}
