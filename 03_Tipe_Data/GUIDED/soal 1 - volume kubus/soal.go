package main

import "fmt"

func main() {
	var sisi, volume int
	fmt.Print("Masukan sisi = ")
	fmt.Scan(&sisi)
	volume = sisi * sisi * sisi
	fmt.Println("Volume kubus = ", volume)
}
