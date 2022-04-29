package main

import (
	"fmt"
	"time"
)

func main() {
	var nama_pengguna string
	var dna_pengguna string
	var nama_penyakit string
	var metode int
	var hasil string

	currentTime := time.Now()
	today := currentTime.Format("2006-01-02")

	fmt.Print("Masukkan nama pengguna: ")
	fmt.Scanln(&nama_pengguna)

	fmt.Print("Masukkan sequence DNA pengguna: ")
	fmt.Scanln(&dna_pengguna) // C:\Users\johan\OneDrive\Documents\GitHub\Tubes3_13520063\test\pengguna\

	fmt.Print("Masukkan nama penyakit: ")
	fmt.Scanln(&nama_penyakit) // C:\Users\johan\OneDrive\Documents\GitHub\Tubes3_13520063\test\penyakit\

	fmt.Println("1. KMP Algorithm")
	fmt.Println("2. Boyer-Moore Algorithm")
	fmt.Print("Masukkan metode: ")
	fmt.Scanln(&metode)

	if checkSequence(dna_pengguna) {
		if metode == 1 {
			if bmMatch(readFile(dna_pengguna), readFile(nama_penyakit)) == -1 {
				hasil += "False"
			} else {
				hasil += "True"
			}
		} else if metode == 2 {
			if kmpMatch(readFile(dna_pengguna), readFile(nama_penyakit)) == -1 {
				hasil += "False"
			} else {
				hasil += "True"
			}
		}
		fmt.Printf("Hasil tes: %s - %s - %s - %s", today, nama_pengguna, nama_penyakit, hasil)
	} else {
		fmt.Println("Masukan sequence DNA pengguna tidak valid")
	}
}
