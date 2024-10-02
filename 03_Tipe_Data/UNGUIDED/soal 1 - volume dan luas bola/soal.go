package main

import "fmt"

func main() {
	var r, volume, luas float64
	const pi = 3.1415926535
	fmt.Scan(&r)
	volume = (4.0 / 3.0) * pi * r * r * r
	luas = 4 * pi * r * r
	fmt.Println("Bola dengan jejari", r, "memiliki volume", volume, "dan", "luas kulit", luas)
}
