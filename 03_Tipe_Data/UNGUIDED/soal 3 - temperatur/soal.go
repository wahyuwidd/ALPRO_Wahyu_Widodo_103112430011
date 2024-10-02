package main

import "fmt"

func main() {
	var celsius, fahrenheit, reamur, kelvin float64
	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius)
	fahrenheit = (celsius * 9 / 5) + 32
	reamur = (celsius * 4 / 5)
	kelvin = (fahrenheit + 459.67) * 5 / 9
	fmt.Println("Derajat Reamur:", reamur)
	fmt.Println("Derajat Fahrenheit:", fahrenheit)
	fmt.Println("Derajat Kelvin:", int(kelvin))
}
