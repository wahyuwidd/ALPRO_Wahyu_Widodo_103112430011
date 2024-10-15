package main

import "fmt"

func main() {
	var n, i int
	fmt.Print("Input bilangan: ")
	fmt.Scan(&n)

	if n <= 0 { //agar input user harus bilangan bulat positif
		fmt.Println("Input harus bilangan bulat positif.")
		return
	}

	sum := 0                 // inisialisasi sum/hasil penjumlahan dengan nilai 0
	for i = 1; i <= n; i++ { // jika i <= n, maka iterasi ++ akan berhenti ketika i > n
		sum += i //cara lain yang singkat dari sum = sum + i
	}

	fmt.Printf("Jumlah dari 1 hingga %d adalah %d\n", n, sum)
}
