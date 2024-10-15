package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, x3 float64
	var y1, y2, y3 float64
	var sisiPanjang float64

	
	fmt.Scanln(&x1, &y1)
	fmt.Scanln(&x2, &y2)
	fmt.Scan(&x3, &y3)
	var AB float64 = math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	var BC float64 = math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	var AC float64 = math.Sqrt(math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2))

	sisiPanjang = math.Max(AB, math.Max(BC, AC))
	fmt.Println(sisiPanjang)
	fmt.Print(x1, x2, x3)
}
