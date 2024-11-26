package main

import "fmt"

func main() {
	var input, jam int
	var teks string
	fmt.Scan(&input)
	switch {
	case input == 0:
		jam = 12
		teks = "AM"
	case input > 0 && input < 12:
		jam = input
		teks = "AM"
	case input == 12:
		jam = 12
		teks = "PM"
	case input > 12 && input <= 23:
		jam = input - 12
		teks = "PM"
	default:
		fmt.Print("Jam tidak valid")
		return
	}

	fmt.Println(jam, teks)
}
