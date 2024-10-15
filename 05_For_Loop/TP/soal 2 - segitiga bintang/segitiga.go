package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input bilangan: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ { // iterasi loop pertama untuk membuat baris segitiga
		for j := 1; j <= i; j++ { // iterasi loop kedua untuk membuat kolom segitiga
			fmt.Print("*")
		}
		fmt.Println() // agar membuat baris baru jadi * nya kebawah/vertikal
	}
}
