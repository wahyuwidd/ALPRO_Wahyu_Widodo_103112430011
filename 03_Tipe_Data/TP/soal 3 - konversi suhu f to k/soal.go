package main

import "fmt"

func main() {
	var f, k float64 //f untuk derajat fahrenheit dan k untuk derajat kelvin
	fmt.Scan(&f)
	k = (f-32)*5/9 + 273.15 //rumus untuk konversi fahrenheit ke kelvin
	fmt.Println(k)          //menampilkan hasil konversi
}
