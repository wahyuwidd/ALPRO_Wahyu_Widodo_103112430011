package main

import "fmt"

func main() {
	var input int
	var continueloop bool
	for continueloop = true; continueloop; {
		fmt.Print("Masukkan bilangan: ")
		fmt.Scan(&input)
		continueloop = input <= 0
	}
	fmt.Println(input, "adalah bilangan bulat positif")
}
