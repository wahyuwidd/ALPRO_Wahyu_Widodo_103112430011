package main

import (
	"fmt"
)

func main() {
	var bmi, tinggiBadan, beratBadan float64
	fmt.Scan(&bmi, &tinggiBadan)
	beratBadan = bmi * tinggiBadan * tinggiBadan //rumus hitung bb dari nilai bmi dan tinggi badan
	fmt.Println(int(beratBadan + 1))             //karena menggunakan int di bulatkan kebawah jadi di tambah dengan 1
}
