package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if nilai > 90 {
		//jika nilai yang di input user lebih dari 90 maka print A
		fmt.Print("A")
	} else if nilai >= 80 && nilai <= 90 {
		//jika nilai yang di input user lebih dari sama dengan 80 dan kurang dari sama dengan 90 maka print AB
		fmt.Print("AB")
	} else if nilai >= 70 && nilai <= 80 {
		//jika nilai yang di input user lebih dari sama dengan 70 dan kurang dari sama dengan 80 maka print B
		fmt.Print("B")
	} else {
		//jika semua kondisi diatas tidak terpenuhi maka print C
		fmt.Print("C")
	}
}
