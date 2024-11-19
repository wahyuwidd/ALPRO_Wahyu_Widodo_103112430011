package main

import "fmt"

func main() {
	var berat int         // berat parsel dalam gram
	var totalBeratKg int  // total berat dalam kg
	var sisaBeratGram int // sisa berat dalam gram
	var biayaKg int       // biaya untuk berat dalam kg
	var biayaSisa int     // biaya untuk sisa berat
	var totalBiaya int    // total biaya pengiriman

	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&berat)

	// Menghitung total berat dalam kg dan sisa berat dalam gram
	totalBeratKg = berat / 1000
	sisaBeratGram = berat % 1000

	// Menghitung biaya pengiriman untuk berat dalam kg
	biayaKg = totalBeratKg * 10000 // Rp. 10.000 per kg

	// Menghitung biaya untuk sisa berat
	if totalBeratKg > 10 {
		// Jika total berat lebih dari 10 kg, sisa berat digratiskan
		biayaSisa = 0
	} else {
		if sisaBeratGram >= 500 {
			biayaSisa = sisaBeratGram * 5 // Rp. 5 per gram
		} else {
			biayaSisa = sisaBeratGram * 15 // Rp. 15 per gram
		}
	}

	// Menghitung total biaya
	totalBiaya = biayaKg + biayaSisa

	// Menampilkan hasil
	fmt.Printf("Detail berat: %d kg + %d gr\n", totalBeratKg, sisaBeratGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
