package main

import "fmt"

func main() {
	var tebakan, jawaban, percobaan, Maxpercobaan int8 //pake int8 karena nilai nya tidak melebihi 100
	jawaban = 39                                       //contoh jawaban angka yang benar 39
	Maxpercobaan = 5                                   //maksimal percobaan

	//jika percobaan lebih dari sama dengan maksimal percobaan maka akan berhenti
	for percobaan = 1; percobaan <= Maxpercobaan; percobaan++ {
		fmt.Print("Tebak angka antara 1 hingga 100: ")
		fmt.Scan(&tebakan)

		if tebakan < 1 || tebakan > 100 { //validasi tebakan hanya antara 1 hingga 100
			fmt.Println("Tebakan harus antara 1 hingga 100")
			return
		}

		//menggunakan percabangan if else untuk menentukan apakah tebakan lebih kecil atau lebih besar atau benar
		if tebakan < jawaban {
			fmt.Println("Tebakan terlalu kecil")
		} else if tebakan > jawaban {
			fmt.Println("Tebakan terlalu besar")
		} else {
			fmt.Println("Selamat, tebakan kamu benar")
			return
		}
	}

	fmt.Printf("Maaf, kamu gagal menebak angkanya. Angka yang benar adalah %d.\n", jawaban)

}
