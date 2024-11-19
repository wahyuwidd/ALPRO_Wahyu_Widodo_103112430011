package main

import "fmt"

func main() {
	var input int
	var menu string
	//menampilkan daftar menu ke user
	fmt.Println("Menu Restoran Cepat Saji:")
	fmt.Println("1. Burger - Rp25,000")
	fmt.Println("2. Fried Chicken - Rp20,000")
	fmt.Println("3. French Fries - Rp15,000")
	fmt.Println("4. Soft Drink - Rp10,000")
	fmt.Println("5. Coffee - Rp15,000")
	//program meminta input dari user yaitu kode menu
	fmt.Print("Masukkan kode menu (1-5): ")
	fmt.Scan(&input)
	//kemudian program melakukan pernyataan
	switch input {
	case 1:
		menu = "Burger - Rp25,000"
	case 2:
		menu = "Fried Chicken - Rp20,000"
	case 3:
		menu = "French Fries - Rp15,000"
	case 4:
		menu = "Soft Drink - Rp10,000"
	case 5:
		menu = ". Coffee - Rp15,000"
	default:
		fmt.Print("Kode menu tidak valid")
	}
	fmt.Print("Anda memilih ", menu)
}
