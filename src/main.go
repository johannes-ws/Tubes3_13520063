package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

func main() {
	var filename string

	fmt.Println("Masukkan nama file: ")
	fmt.Scanln(&filename)
	checkSequence(filename)
}
