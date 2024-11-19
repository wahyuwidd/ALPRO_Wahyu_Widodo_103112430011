package main

import "fmt"

func main() {
	var b int
	var p bool

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)
	jumlah := 0
	if b <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1!")
	} else {
		fmt.Printf("Faktor dari %d adalah: ", b)
		for i := 1; i <= b; i++ {
			if b%i == 0 {
				fmt.Printf("%d ", i)
				jumlah++
			}
		}
		fmt.Println()
		if jumlah == 2 {
			p = true
		} else {
			p = false
		}
		fmt.Print("Prima: ", p)
	}
}
