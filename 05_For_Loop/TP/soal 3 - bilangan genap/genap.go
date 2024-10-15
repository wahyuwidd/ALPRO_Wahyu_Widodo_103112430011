package main

import "fmt"

func main() {
	var nilaiMax = 50
	for i := 1; i <= nilaiMax; i++ { //jika i >= nilaiMax maka iterasi akan berhenti berarti akan cetak antara 1 hingga nilaiMax
		if i%2 == 0 { //sebelum di cetak lakukan pernyataan jika i habis dibagi 2 maka bilangan itu genap
			fmt.Print(i, " ") //jika pernyataan benar maka yang dicetak hanya bilangan genap
			//kemudian dipisahkan dengan spasi menggunakan " " string kosong
		}
	}
}
