package main

import (
	"fmt"
)

func main() {
	var keliling, luas int
	const sisi = 27     //karena nilai panjang sisi persegi alun" pwt 27 maka harus menggunakan konstanta biar nilai nya tetap.
	keliling = 4 * sisi //rumus menghitung keliling persegi
	luas = sisi * sisi  //rumus menghitung luas persegi

	fmt.Println("Keliling Persegi:", keliling, "cm")
	fmt.Println("Luas Persegi:", luas, "cm")
}
