package main

import "fmt"

func main() {
	var umur int
	var kewarganegaraan string
	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umur)
	fmt.Print("Masukkan kewarganegaraan: ")
	fmt.Scan(&kewarganegaraan)
	if umur < 17 && kewarganegaraan != "WNI" {
		//jika inputan user umur kurang dari 17 dan kewarganegaraan bukan WNI maka print "Anda masih dibawah umur dan bukan WNI tidak bisa mengikuti pemilu"
		fmt.Print("Anda masih dibawah umur dan bukan WNI tidak bisa mengikuti pemilu")
	} else if umur < 17 {
		//jika inputan user umur kurang dari 17 maka print "Maaf anda masih di bawah umur tidak bisa mengikuti pemilu"
		fmt.Print("Maaf anda masih di bawah umur tidak bisa mengikuti pemilu")
	} else if kewarganegaraan != "WNI" {
		//jika inputan user kewarganegaraan bukan WNI maka print "Maaf anda bukan WNI tidak bisa mengikuti pemilu"
		fmt.Print("Maaf anda bukan WNI tidak bisa mengikuti pemilu")
	} else {
		//jika semua kondisi diatas tidak dipenuh maka print "Anda bisa mengikuti pemilu" atau
		//Jika inputan user umur lebih dari 17 dan kewarganegaraan WNI maka print "Anda bisa mengikuti pemilu"
		fmt.Print("Anda bisa mengikuti pemilu")
	}
}
