package main

import "fmt"

func main() {
	var input float64
	var selesai bool
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scanln(&input)

	target := int(input)
	if input > float64(target) {
		target++
	}

	jumlah := input
	step := 0.1

	for selesai = false; !selesai; {
		fmt.Printf("%.1f\n", jumlah)
		jumlah += step
		if jumlah >= float64(target) {
			selesai = true
			fmt.Print(target)
		}
	}
}
