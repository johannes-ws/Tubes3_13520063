package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

func checkSequence(filename string) bool {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("^[ACTG]+$")
	return re.Match(content)
}

func main() {
	var nama_pengguna string
	var dna_pengguna string
	var nama_penyakit string
	var metode int
	var hasil string

	currentTime := time.Now()
	today := currentTime.Format("02 January 2006")

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
			if kmp(readFile(dna_pengguna), readFile(nama_penyakit)) == -1 {
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
