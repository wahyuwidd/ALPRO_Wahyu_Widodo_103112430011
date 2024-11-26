package main

import "fmt"

func main() {
	var nama_barang string
	var opsi, harga_barang, totalBelanja int
	fmt.Println("Selamat datang di Aplikasi Kasir Sederhana!")
	fmt.Println("===========================================")
	for { //agar selalu melakukan  loop sampai user memilih opsi 2(terdapat return untuk menghentikan perulangan)
		//terdapat 2 menu yg bisa di pilih user
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambahkan barang")
		fmt.Println("2. Selesaikan transaksi")
		fmt.Print("Pilih opsi (1/2): ")
		fmt.Scan(&opsi)
		switch opsi {
		case 1: //jika user menginputkan opsi 1 maka program meminta user menginputkan kembali yaitu nama barang dan harga barang nya
			fmt.Print("Masukkan nama barang: ")
			fmt.Scan(&nama_barang)
			fmt.Print("Masukkan harga barang: ")
			fmt.Scan(&harga_barang)
			//kemudian jumlahkan harga barang yg di input user ke dalam variable total belanja
			totalBelanja += harga_barang
			//menampilkan detail barang dan harga barang yang berhasil di tambah
			fmt.Println("Barang", nama_barang, "dengan harga Rp", harga_barang, "berhasil ditambahkan")
		case 2:
			//jika user menginputkan opsi 2 maka program langsung manampilkan total belanja yang sudah dihitung di opsi 1
			fmt.Println("\n===========================================")
			fmt.Println("Total belanja Anda: ", totalBelanja)
			fmt.Println("Terima kasih telah menggunakan aplikasi kasir!")
			fmt.Println("===========================================")
			return //untuk keluar dari perulangan
		default:
			fmt.Println("Opsi tidak valid. Silakan pilih 1 atau 2.")
		}
	}
}
