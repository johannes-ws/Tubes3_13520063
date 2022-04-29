# Tugas Besar 3 IF2211 Strategi Algoritma
# Penerapan String Matching dan Regular Expression dalam DNA Pattern Matching

|             Nama            |    NIM   |
| --------------------------- | :------: |
| Louis Yanggara              | 13520063 |
| Johannes Winson Sukiatmodjo | 13520123 |
| Daffa Romyz Aufa            | 13520162 |

## Deskripsi Singkat Program
Program ini merupakan sebuah aplikasi berbasis website yang dapat mendeteksi apakah seorang pasien mempunyai penyakit genetik tertentu serta menyimpan hasil prediksi tersebut pada basis data untuk kemudian dapat ditampilkan berdasarkan query pencarian. Dengan memanfaatkan algoritma _String Matching_ dan _Regular Expression_, program ini akan mengecek kevalidan dari masukan sequence DNA, melakukan pencocokan DNA pengguna dengan DNA penyakit, serta menampilkan data-data hasil prediksi tersebut berdasarkan query pencarian. Program ini terdiri dari 3 menu utama, yaitu menu `Tambahkan Penyakit`, menu `Tes DNA`, dan menu `Pencarian Hasil Prediksi`. Pada menu `Tambahkan Penyakit`, program ini menerima input nama penyakit dan sequence DNA-nya. Pada menu `Tes DNA`, program ini menerima input nama pengguna, sequence DNA-nya, dan prediksi penyakitnya serta mengeluarkan output berupa hasil tesnya. Pada menu `Pencarian Hasil Prediksi`, program ini menerima input query pencarian serta mengeluarkan output berupa data-data yang sesuai dengan query tersebut. Program ini dibuat dengan Backend menggunakan bahasa Go dan Frontend menggunakan React.

## Requirement Program dan Instalasi _Module_/_Package_
1. Visual Studio Code --> https://code.visualstudio.com/download
2. Go --> https://go.dev/doc/install
3. React --> https://docs.microsoft.com/en-us/windows/dev-environment/javascript/react-on-windows
4. MySQL --> https://www.mysql.com/downloads/

## Cara Menggunakan Program
1. Clone repository ini
2. Buka Visual Studio Code
3. Buka folder repository ini
4. Buka terminal baru
5. Ketik `cd src/backend`
6. Ketik `go build`
7. Buka terminal baru kembali
8. Ketik `cd src/frontend/dna`
9. Ketik `npm start`