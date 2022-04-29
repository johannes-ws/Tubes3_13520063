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
	var pilihan int

	fmt.Println("1. Insert penyakit")
	fmt.Println("2. Memprediksi penyakit seseorang")
	fmt.Println("3. Search hasil tes")
	fmt.Println("4. Exit")
	fmt.Print("Fitur yang ingin digunakan: ")
	fmt.Scanln(&pilihan)

	for pilihan != 4 {
		if pilihan == 1 {
			var nama string
			var filename string
			fmt.Printf("Masukkan nama penyakit: ")
			fmt.Scanln(&nama)
			fmt.Printf("masukkan file sequence: ")
			fmt.Scanln(&filename)
			path := "\\Tubes3_13520063\\test\\penyakit\\" + filename
			if checkSequence(path) {
				sequence := readFile(path)
				InsertPenyakit(nama, sequence)
			} else {
				fmt.Println("Masukkan sequence salah!")
			}

		} else if pilihan == 2 {
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

			dna_path := "\\Tubes3_13520063\\test\\pengguna\\" + dna_pengguna
			nama_path := "\\Tubes3_13520063\\test\\penyakit\\" + nama_penyakit
			if checkSequence(dna_path) {
				if metode == 1 {
					if bmMatch(readFile(dna_path), readFile(nama_path)) == -1 {
						hasil += "False"
					} else {
						hasil += "True"
					}
				} else if metode == 2 {
					if kmpMatch(readFile(dna_path), readFile(nama_path)) == -1 {
						hasil += "False"
					} else {
						hasil += "True"
					}
				}
				fmt.Printf("Hasil tes: %s - %s - %s - %s\n", today, nama_pengguna, nama_penyakit, hasil)
				InsertHasilTes(today, nama_pengguna, nama_penyakit, hasil)
			} else {
				fmt.Println("Masukan sequence DNA pengguna tidak valid")
			}
		} else if pilihan == 3 {
			var query string
			fmt.Print("Masukkan query pencarian (YYYY-MM-DD<spasi>Nama_Penyakit atau YYYY-MM-DD atau Nama_Penyakit): ")
			fmt.Scanln(&query)
			checkSearchInput(query)
		}
		fmt.Println("1. Insert penyakit")
		fmt.Println("2. Memprediksi penyakit seseorang")
		fmt.Println("3. Search hasil tes")
		fmt.Println("4. Exit")
		fmt.Print("Fitur yang ingin digunakan: ")
		fmt.Scanln(&pilihan)
	}

}
