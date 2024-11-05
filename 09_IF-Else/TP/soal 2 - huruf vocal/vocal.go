package main

import "fmt"

func main() {
	var input string
	fmt.Print("Input huruf: ")
	fmt.Scan(&input)
	if input == "A" || input == "I" || input == "U" || input == "E" || input == "O" {
		//jika inputan user adalah A, I, U, E, O maka print Huruf Vokal
		fmt.Print("Huruf Vokal")
	} else {
		//jika bukan A, I, U, E, O maka print Huruf Konsonan
		fmt.Print("Huruf Konsonan")
	}
}
