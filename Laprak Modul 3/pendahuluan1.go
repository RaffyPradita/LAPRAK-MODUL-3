package main

import "fmt"

func main1() {
	var N, Sisa, Jam, Menit, Detik int // Deklarasi variabel

	fmt.Print("Masukkan N : ") // Meminta input dari pengguna

	fmt.Scan(&N) // Membaca input jumlah waktu dalam detik

	Jam = N / 3600    // menghitung jumlah jam
	Sisa = N % 3600   // menghitung sisa detik setelah jam
	Menit = Sisa / 60 // menghitung jumlah menit
	Detik = Sisa % 60 // menghitung sisa detik setelah menit

	fmt.Println(Jam, Menit, Detik) // Menampilkan hasil konversi waktu dalam format jam, menit, dan detik
}
