package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type hasil struct {
	id            int
	tanggal       string
	nama_pasien   string
	nama_penyakit string
	status        string
}

const ( // isi sesuai punya kalian
	db_user     = "root"
	db_password = "Johannes_235711"
	db_host     = "localhost"
	db_port     = 1234
	db_database = "dna"
)

func InsertPenyakit(nama, sequence string) {
	s := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db_user, db_password, db_host, db_port, db_database)
	db, err := sql.Open("mysql", s)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO penyakit (nama_penyakit, sequence) VALUES ('%s','%s');", nama, sequence)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func checkSearchInput(input string) {
	s := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db_user, db_password, db_host, db_port, db_database)
	db, err := sql.Open("mysql", s)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	re1 := regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])`)
	re2 := regexp.MustCompile(`[A-Z][a-zA-Z]*(\s*[A-Z][a-zA-Z]*)*`)
	var date []string
	var name []string
	if re1.MatchString(input) && re2.MatchString(input) {
		fmt.Println("Pencarian berdasarkan tanggal dan nama")
		date = re1.FindAllString(input, -1)
		name = re2.FindAllString(input, -1)
		namaReal := strings.Join(name, "")
		tanggalReal := strings.Join(date, "")
		rows, err := db.Query("select id, tanggal, nama_pasien, nama_penyakit, status from test where tanggal = '%s' and nama_penyakit = '%s'", tanggalReal, namaReal)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()

		var result []hasil

		for rows.Next() {
			var each = hasil{}
			var err = rows.Scan(&each.id, &each.tanggal, &each.nama_pasien, &each.nama_penyakit, &each.status)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result = append(result, each)
		}

		if err = rows.Err(); err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, each := range result {
			fmt.Println(each.id, each.tanggal, each.nama_pasien, each.nama_penyakit, each.status)
		}
	} else if re1.MatchString(input) {
		fmt.Println("Pencarian berdasarkan tanggal saja")
		date = re1.FindAllString(input, -1)
		tanggalReal := strings.Join(date, "")
		rows, err := db.Query("select id, tanggal, nama_pasien, nama_penyakit, status from test where tanggal = ?", tanggalReal)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()

		var result []hasil

		for rows.Next() {
			var each = hasil{}
			var err = rows.Scan(&each.id, &each.tanggal, &each.nama_pasien, &each.nama_penyakit, &each.status)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result = append(result, each)
		}

		if err = rows.Err(); err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, each := range result {
			fmt.Println(each.id, each.tanggal, each.nama_pasien, each.nama_penyakit, each.status)
		}
	} else if re2.MatchString(input) {
		fmt.Println("Pencarian berdasarkan nama saja")
		name = re2.FindAllString(input, -1)
		namaReal := strings.Join(name, "")
		rows, err := db.Query("select id, tanggal, nama_pasien, nama_penyakit, status from test where nama_penyakit = ?", namaReal)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()

		var result []hasil

		for rows.Next() {
			var each = hasil{}
			var err = rows.Scan(&each.id, &each.tanggal, &each.nama_pasien, &each.nama_penyakit, &each.status)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result = append(result, each)
		}

		if err = rows.Err(); err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, each := range result {
			fmt.Println(each.id, each.tanggal, each.nama_pasien, each.nama_penyakit, each.status)
		}
	} else {
		fmt.Println("Format tidak sesuai (YYY-MM-DD) NamaPenyakit")
	}

}

func readFile(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func checkSequence(filename string) bool {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("^[ACTG]+$")
	return re.Match(content)
}

func InsertHasilTes(today, nama_pengguna, nama_penyakit, hasil string) {
	s := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db_user, db_password, db_host, db_port, db_database)
	db, err := sql.Open("mysql", s)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO test (tanggal, nama_pasien, nama_penyakit, status) VALUES ('%s','%s','%s','%s')", today, nama_pengguna, nama_penyakit, hasil)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

// func main() {
// 	penyakit1 := "HIV"
// 	sequence := "ACTGCTGCTGCAT"
// 	InsertPenyakit(penyakit1, sequence)
// }
