package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func checkSequence(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	re := regexp.MustCompile("^[ACTG]+$")
	fmt.Println(re.Match(content))
}

func checkSearchInput(input string) (string, string) {
	re1 := regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])`)
	re2 := regexp.MustCompile(`[A-Z][a-zA-Z]*(\s*[A-Z][a-zA-Z]*)*`)
	var date []string
	var name []string
	if re1.MatchString(input) && re2.MatchString(input) {
		fmt.Println("Pencarian berdasarkan tanggal dan nama")
		date = re1.FindAllString(input, -1)
		name = re2.FindAllString(input, -1)
	} else if re1.MatchString(input) {
		fmt.Println("Pencarian berdasarkan tanggal saja")
		date = re1.FindAllString(input, -1)
	} else if re2.MatchString(input) {
		fmt.Println("Pencarian berdasarkan nama saja")
		name = re2.FindAllString(input, -1)
	} else {
		fmt.Println("Format tidak sesuai (YYY-MM-DD) NamaPenyakit")
	}

	namaReal := strings.Join(name, "")
	tanggalReal := strings.Join(date, "")
	return namaReal, tanggalReal
}

func main() {
	var filename string

	fmt.Print("Masukkan nama file: ")
	fmt.Scanln(&filename)
	checkSequence(filename)
	str4 := "2022-11-11 Sakit Gigi"
	nama, tanggal := checkSearchInput(str4)
	fmt.Println(nama, tanggal)
}
