package main

import "fmt"

func main() {
	var angka1, angka2, hasil float64
	var operasi string
	fmt.Print("Masukan angka 1: ")
	fmt.Scanln(&angka1)
	fmt.Print("Pilih Operasi Aritmatika (+, -, *, /): ") //operasi aritmatika
	fmt.Scanln(&operasi)
	fmt.Print("Masukan angka 2: ")
	fmt.Scan(&angka2)

	switch operasi {
	case "+":
		hasil = angka1 + angka2
	case "-":
		hasil = angka1 - angka2
	case "*":
		hasil = angka1 * angka2
	case "/":
		if angka2 != 0 {
			hasil = angka1 / angka2
		} else {
			fmt.Println("Error: Tidak bisa membagi dengan nol")
		}
	default:
		fmt.Println("Operasi tidak valid. Gunakan +, -, *, atau /")
	}
	fmt.Println("Hasil dari", angka1, operasi, angka2, "adalah", hasil)

}
