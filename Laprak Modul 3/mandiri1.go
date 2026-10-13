package main

import "fmt"

func main4() {
	var emas, sisa, perak, tembaga, koin int // Deklarasi variabel
	fmt.Print("Masukkan jumlah koin : ")     // Meminta input dari pengguna
	fmt.Scan(&koin)                          // Membaca input jumlah koin

	emas = koin / 9    // menghitung jumlah emas
	sisa = koin % 9    // menghitung sisa koin setelah menghitung emas
	perak = sisa / 3   // menghitung jumlah perak
	tembaga = koin % 3 // menghitung jumlah tembaga

	fmt.Println(emas, perak, tembaga) // Menampilkan hasil konversi koin
}
