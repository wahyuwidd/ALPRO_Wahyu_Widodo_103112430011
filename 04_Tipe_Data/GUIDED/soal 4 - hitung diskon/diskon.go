package main

import (
	"fmt"
)

func main() {
	var totalBelanjaAwal, totalBelanjaAkhir, diskon int
	fmt.Scanln(&totalBelanjaAwal)
	fmt.Scanln(&diskon)
	totalBelanjaAkhir = totalBelanjaAwal - (totalBelanjaAwal * diskon / 100) //rumus menghitung diskon
	fmt.Println("Total belanja setelah diskon: ", totalBelanjaAkhir)
}
