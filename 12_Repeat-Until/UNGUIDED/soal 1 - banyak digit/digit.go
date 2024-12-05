package main

import "fmt"

func main() {
	var input, total int
	var selesai bool
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&input)
	for selesai = false; !selesai; {
		total++
		input /= 10
		selesai = input == 0
	}
	fmt.Println(total)
}
