package main

import "fmt"

func main() {
	var a, i int
	var faktorial int
	fmt.Scan(&a)
	faktorial = 1
	for i = 1; i <= a; i++ {
		faktorial *= i
	}
	fmt.Println(faktorial)
}
