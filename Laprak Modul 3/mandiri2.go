package main

import "fmt"

func main() {
	var Gp, lembur, Bl, Pt, Gb int // Deklarasi variabel

	fmt.Print("Masukkan gaji pokok dan jumlah lembur : ") // Meminta input dari pengguna
	fmt.Scan(&Gp, &lembur)                                // Membaca input gaji pokok dan jumlah lembur

	Bl = lembur * 45000 // Menghitung bonus lembur
	Pt = Gp * 55 / 1000 // Menghitung total potongan
	Gb = Gp + Bl - Pt   // Menghitung gaji bersih

	fmt.Println(Gb) // Menampilkan gaji bersih
}
