# Order API Golang

## Tujuan

Tujuan dari proyek ini adalah untuk membangun REST API di Golang yang dapat digunakan untuk mengelola pesanan.

## Fitur

Proyek ini memiliki beberapa fitur utama:

- Membuat pesanan baru (endpoint: `/orders`)
- Memperbarui detail pesanan yang sudah ada (endpoint: `/orders/:orderID`)
- Menghapus pesanan (endpoint: `/orders/:orderID`)
- Mengambil detail pesanan (endpoint: `/orders/:orderID`)

## Data Model

Data model yang digunakan dalam proyek ini terdiri dari beberapa atribut utama:

- ID pesanan
- Nama pelanggan
- Daftar item yang dipesan (termasuk kode item, deskripsi, dan jumlah)
- Status pesanan

## Error Handling

Proyek ini dilengkapi dengan penanganan error yang memadai, dengan beberapa aspek sebagai berikut:

- Menampilkan pesan error yang jelas untuk setiap jenis error.
- Mencatat error di log file.

## Library atau Framework

Proyek ini menggunakan beberapa library atau framework, antara lain:

- Gin Gonic: Digunakan untuk routing dan middleware.
- GORM: Digunakan untuk interaksi dengan database.

