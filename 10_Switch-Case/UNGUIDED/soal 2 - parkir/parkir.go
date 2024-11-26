package main

import "fmt"

func main() {
	var jenis_kendaraan string
	var durasi_parkir, tarif int
	fmt.Scan(&jenis_kendaraan, &durasi_parkir)
	switch jenis_kendaraan {
	case "motor":
		tarif = 2000 * durasi_parkir
	case "mobil":
		tarif = 5000 * durasi_parkir
	case "truk":
		tarif = 8000 * durasi_parkir
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}
	fmt.Println("Rp", tarif)
}
