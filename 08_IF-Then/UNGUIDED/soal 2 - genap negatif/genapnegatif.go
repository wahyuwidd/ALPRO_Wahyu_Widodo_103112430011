package main

import "fmt"

func main() {
	var bilangan int
	var teks = "bukan"
	fmt.Scan(&bilangan)
	if bilangan < 0 && bilangan%2 == 0 {
		teks = "genap negatif"
	}
	fmt.Print(teks)
}
