package main

import "fmt"

func main() {
	var bilangan int
	var teks = "Bukan Positif"
	fmt.Scan(&bilangan)
	if bilangan > 0 {
		teks = "Positif"
	}
	fmt.Print(teks)
}
