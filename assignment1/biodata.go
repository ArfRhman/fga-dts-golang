package main

import (
	"fmt"
	"os"
)

// Struct untuk menyimpan data peserta
type Peserta struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Fungsi untuk mendapatkan data peserta berdasarkan nomor absen
func getPesertaByAbsen(absen int) Peserta {
	peserta := map[int]Peserta{
		1: {"John Doe", "Jalan Raya 123", "Software Engineer", "Ingin belajar pemrograman"},
		2: {"Jane Smith", "Jalan Melati 456", "Data Analyst", "Mengembangkan keterampilan baru"},
		3: {"Mark Johnson", "Jalan Mawar 789", "Graphic Designer", "Pengalaman di industri kreatif"},
		4: {"Emily Brown", "Jalan Anggrek 101", "Marketing Manager", "Ingin memperluas jaringan profesional"},
		5: {"Michael Lee", "Jalan Teratai 202", "Teacher", "Pengalaman mengajar dan ingin mendalami teknologi"},
	}

	// Mengembalikan data peserta sesuai nomor absen
	return peserta[absen]
}

func main() {
	// Mendapatkan argumen dari command line
	args := os.Args

	// Mengonversi argumen ke dalam bentuk integer
	absen := 0
	fmt.Sscanf(args[1], "%d", &absen)

	// Mendapatkan data peserta berdasarkan nomor absen
	peserta := getPesertaByAbsen(absen)

	// Menampilkan data peserta
	fmt.Println("Nama:", peserta.Nama)
	fmt.Println("Alamat:", peserta.Alamat)
	fmt.Println("Pekerjaan:", peserta.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", peserta.Alasan)
}
