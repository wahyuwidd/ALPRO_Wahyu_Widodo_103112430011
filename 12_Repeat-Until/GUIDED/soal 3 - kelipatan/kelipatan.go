package main

import "fmt"

func main() {
	var x, y int
	var selesai bool
	fmt.Print("Masukkan bilangan x dan y (x >= y): ")
	fmt.Scan(&x, &y)
	for selesai = false; !selesai; {
		x -= y
		selesai = x <= 0
		fmt.Println(x)
	}
	fmt.Print(x == 0)
}
