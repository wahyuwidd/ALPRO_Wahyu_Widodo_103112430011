package main

import (
	"fmt"
)

func main() {
	var x int                 //karena variabel x bilangan bulat pake tipe data integer
	var fx float64            //karena hasil fx bisa jadi desimal maka tipe data nya pake bilangan real atau float64
	fmt.Scan(&x)              //untuk menerima input dari user dan menyimpan hasil input di variabel x
	fx = 2.0/float64(x+5) + 5 //rumus persaman f(x)
	fmt.Println(fx)           //menampilkan output
}
