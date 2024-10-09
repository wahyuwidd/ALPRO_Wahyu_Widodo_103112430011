package main

import (
	"fmt"
)

func main() {
	var detik, jam, menit, masukan int
	fmt.Scan(&masukan)
	jam = masukan / 3600
	menit = (masukan % 3600) / 60
	detik = masukan % 60
	fmt.Println(jam, "jam", menit, "menit dan", detik, "detik")
}
