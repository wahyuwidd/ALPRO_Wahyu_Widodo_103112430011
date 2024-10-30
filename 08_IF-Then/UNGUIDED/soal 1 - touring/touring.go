package main

import "fmt"

func main() {
	var bilangan int
	var motor int
	fmt.Scan(&bilangan)
	motor = bilangan / 2
	if bilangan%2 != 0 {
		motor = bilangan/2 + 1
	}
	fmt.Print(motor)
}
