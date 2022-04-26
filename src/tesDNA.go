package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

// kmp.go
/* mengembalikan border function */
func borderFunction(pattern string) []int {

	var fail [1000000]int
	fail[0] = 0
	var m int = len(pattern)
	var i int = 1 // iterator suffiks
	var j int = 0 // iterator prefiks

	for i < m {
		if pattern[j] == pattern[i] {
			fail[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			fail[i] = 0
			i++
		}
	}
	return fail[0:len(pattern)]
}

/*	menerima text dan pattern,
	mengembalikan indeks pattern pertama yang ditemukan di text,
	mengembalikan -1 jika ditak ada pattern dalam text
*/
func kmp(text string, pattern string) int {
	var n int = len(text)
	var m int = len(pattern)

	var fail []int = borderFunction(pattern)

	var i int = 0 // iterator untuk text
	var j int = 0 // iterator untuk pattern

	for i < n {
		if pattern[j] == text[i] {
			if j == m-1 {
				return i - m + 1
			}
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			i++
		}
	}
	return -1
}

// BM.go
func buildLast(pattern string) []int {
	var last [128]int
	for i := 0; i < 128; i++ {
		last[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		last[int(pattern[i])] = i
	}
	return last[:]
}

func bmMatch(text string, pattern string) int {
	last := buildLast(pattern)
	n := len(text)
	m := len(pattern)
	i := m - 1
	if i > n-1 {
		return -1
	}
	j := m - 1
	for {
		if pattern[j] == text[i] {
			if j == 0 {
				return i
			} else {
				i--
				j--
			}
		} else {
			lo := last[int(text[i])]
			i += m - int(math.Min(float64(j), float64(1+lo)))
			j = m - 1
		}
		if i > n-1 {
			break
		}
	}
	return -1
}

func readFile(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
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
}
