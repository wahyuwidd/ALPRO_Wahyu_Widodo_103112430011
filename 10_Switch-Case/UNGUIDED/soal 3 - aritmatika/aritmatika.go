package main

import "fmt"

func main() {
	var input int
	fmt.Print("Masukkan sebuah bilangan bulat: ")
	fmt.Scan(&input)

	switch {
	case input%10 == 0:
		// Bilangan Kelipatan 10
		hasil := input / 10
		fmt.Printf("Kategori: Bilangan Kelipatan 10\n")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", input, hasil)

	case input%5 == 0 && input%2 == 0:
		// Bilangan Kelipatan 5
		hasil := input * input
		fmt.Printf("Kategori: Bilangan Kelipatan 5\n")
		fmt.Printf("Hasil kuadrat dari %d ^ 2 = %d\n", input, hasil)

	case input%2 == 0:
		// Bilangan Genap
		hasil := input * (input + 1)
		fmt.Printf("Kategori: Bilangan Genap\n")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", input, input+1, hasil)

	default:
		// Bilangan Ganjil
		hasil := input + (input + 1)
		fmt.Printf("Kategori: Bilangan Ganjil\n")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", input, input+1, hasil)
	}
}
