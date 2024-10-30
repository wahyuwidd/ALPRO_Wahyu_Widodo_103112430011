package main

import "fmt"

func main() {
	var x, y int
	var faktor1 bool
	var faktor2 bool

	fmt.Scan(&x, &y)
	if y%x == 0 {
		faktor1 = true
	}

	if x%y == 0 {
		faktor2 = true
	}

	fmt.Println(faktor1)
	fmt.Print(faktor2)
}
