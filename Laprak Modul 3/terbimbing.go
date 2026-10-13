package main

import "fmt"

func main3() {
	var kecepatan, menit, jam, total_jarak int // Deklarasi variabel

	fmt.Print("Masukkan kecepatan : ")     // Meminta input dari pengguna
	fmt.Scan(&kecepatan)                   // Membaca input kecepatan dalam km/jam
	total_jarak = 100 + 60 + 170           // menghitung total jarak yang ditempuh dalam km
	menit = (total_jarak * 60) / kecepatan // menghitung waktu tempuh dalam menit

	jam = menit / 60   // menghitung waktu tempuh dalam jam
	menit = menit % 60 // menghitung sisa menit setelah jam

	fmt.Println(jam, menit) // Menampilkan hasil waktu tempuh dalam format jam dan menit
}
