package main

import "fmt"

func main() {
	var angka, total_bil_ganjil int
	fmt.Print("Masukkan bilangan bulat positif n: ")
	fmt.Scan(&angka)
	for i := 1; i <= angka; i++ {
		if i%2 != 0 {
			total_bil_ganjil++
		}
	}
	fmt.Printf("Terdapat %d bilangan ganjil", total_bil_ganjil)
}
