# Enigma Laundry

Enigma Laundry adalah aplikasi sederhana untuk mengelola data customer, service, dan transaction pada usaha laundry.

## Daftar Isi

- [Deskripsi](#deskripsi)
- [Fitur](#fitur)
- [Instalasi](#instalasi)
- [Penggunaan](#penggunaan)


## Deskripsi

Aplikasi Enigma Laundry adalah sebuah perangkat lunak yang memungkinkan pengguna untuk mengelola data customer, service, dan transaction pada usaha laundry. 
Aplikasi ini dikembangkan menggunakan bahasa pemrograman Go dan menggunakan PostgreSQL sebagai basis data.

## Fitur

Aplikasi Enigma Laundry Management System memiliki fitur-fitur berikut:

- Melihat daftar customer
- Menambahkan customer baru
- Mengedit data customer
- Menghapus customer
- Melihat daftar service
- Menambahkan service baru
- Mengedit data service
- Menghapus service
- Melihat daftar transaction
- Menambahkan transaction baru
- Mencetak invoice transaction

## Instalasi

Pastikan Anda telah menginstal Go dan PostgreSQL di komputer Anda sebelum menjalankan aplikasi ini.

1. Clone repositori ini ke dalam direktori lokal Anda.

2. Masuk ke direktori proyek :
cd enigma-laundry


3. Pastikan PostgreSQL berjalan dan sesuaikan konfigurasi koneksi basis data di file `main.go`:
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "enigma_laundry"
)

4. Jalankan perintah berikut untuk menginstal dependensi dan menjalankan aplikasi :
go run main.go


## Penggunaan

Setelah aplikasi berjalan, Anda dapat menggunakan menu yang tersedia pada layar konsol untuk mengakses fitur-fitur yang ada.

1. Saat aplikasi berjalan, pilih nomor menu yang sesuai dengan fitur yang ingin Anda gunakan.
2. Ikuti petunjuk pada layar untuk menambahkan, mengedit, atau menghapus data customer, service, atau transaction.
3. Untuk mencetak invoice transaction, pilih menu "Cetak Invoice" dan masukkan ID customer yang ingin Anda cetak invoice-nya.
