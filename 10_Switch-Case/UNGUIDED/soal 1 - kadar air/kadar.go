package main

import "fmt"

func main() {
	var kadar float64
	var teks string
	fmt.Scan(&kadar)
	switch {
	case kadar >= 6 && kadar <= 8.6:
		teks = "Layak Minum"
	case kadar >= 0 && kadar < 6.5 || kadar > 8.6 && kadar <= 14:
		teks = "Tidak Layak Minum"
	default:
		fmt.Print("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
		return
	}
	fmt.Println("Air", teks)
}
