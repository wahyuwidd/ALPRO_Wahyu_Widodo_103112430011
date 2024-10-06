package main

import (
	"fmt"
)

func main() {
	//variable jam untuk input jam kerja dari user
	//variable upaPerjam untuk input upah per jam dari user
	var jam, upahPerjam, totalGaji, lembur float64
	var jamNormal float64 = 40   //jam kerja normal tanpa lembur
	var bulan float64 = 4        //1 bulan ada 4 minggu
	var upahLembur float64 = 1.5 //upah lembur 1.5 kali upah per jam
	fmt.Print("Masukkan jam kerja: ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan upah per jam: ")
	fmt.Scan(&upahPerjam)
	if jam > jamNormal {
		//menghitung berapa jam lemburnya
		lembur = jam - jamNormal

		//menghitung total gajinya beserta lemburnya dalam sebulan
		totalGaji = (jamNormal*upahPerjam + lembur*upahPerjam*upahLembur) * bulan
	} else {
		//menghitung total gajinya tanpa lembur dalam 4 minggu / sebulan
		totalGaji = jam * upahPerjam * bulan
	}
	//menampilkan total gajinya sekalian di konversi jadi integer yang tadi nya float
	fmt.Print("Total gaji bulanan: Rp.", int(totalGaji))
}
