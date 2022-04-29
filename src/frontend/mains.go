package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var msg string = string(p)
		var a []string = strings.Split(msg, ";")

		currentTime := time.Now()
		today := currentTime.Format("2006-01-02")

		if a[0] == "Tambah" {
			var nama_penyakit string = a[1]
			var sequence string = a[2]
			InsertPenyakit(nama_penyakit, sequence)
		}

		if a[0] == "Test" {
			var nama_pengguna string = a[1]
			var dna_pengguna string = a[2]
			var nama_penyakit string = a[3]
			var hasil string
			if kmpMatch(dna_pengguna, nama_penyakit) == -1 {
				hasil += "False"
			} else {
				hasil += "True"
			}
			InsertHasilTes(today, nama_pengguna, nama_penyakit, hasil)
		}

		if a[0] == "Hasil" {
			var hasiltest []string = checkSearchInput(a[1])
			var bytearray []byte
			for i := 0; i < len(hasiltest); i++ {
				var s = hasiltest[i]
				for j := 0; i < len(s); j++ {
					bytearray = append(bytearray, s[j])
				}
			}
			err = conn.WriteMessage(messageType, bytearray)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "DNA")
	})
	http.HandleFunc("/ws", serveWs)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
