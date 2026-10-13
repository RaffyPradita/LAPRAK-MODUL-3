package main

import "fmt"

func main2() {
	var N, hari int            // Deklarasi variabel
	fmt.Print("Masukkan N : ") // Meminta input dari pengguna
	fmt.Scan(&N)               // Membaca input jumlah hari

	hari = ((4 - 1 + N) % 7) + 1 // menghitung hari setelah N hari dari hari ke-4 (Kamis)

	fmt.Println(hari) // Menampilkan hasil hari setelah N hari dari hari ke-4 (Kamis)
}
