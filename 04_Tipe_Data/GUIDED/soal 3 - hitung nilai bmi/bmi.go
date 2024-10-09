package main

import (
	"fmt"
)

func main() {
	var beratBadan, tinggiBadan, bmi float64
	fmt.Scan(&beratBadan, &tinggiBadan)
	hasil := tinggiBadan * tinggiBadan
	bmi = beratBadan / hasil
	fmt.Printf("%.2f", bmi)
}
