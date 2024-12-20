package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Data produk
type Produk struct {
    Nama   string
    Harga  float64
    Stok   int
}

// Daftar produk
var daftarProduk = []Produk{
    {Nama: "Sabun", Harga: 5000, Stok: 10},
    {Nama: "Minuman", Harga: 3000, Stok: 15},
    {Nama: "Snack", Harga: 2000, Stok: 20},
}

// Fungsi utama
func main() {
    // Menginisialisasi variabel
    var totalBelanja float64 = 0
    var keranjangBelanja []Produk

    // Membuka scanner untuk input pengguna
    scanner := bufio.NewScanner(os.Stdin)

    // Menampilkan daftar produk
    fmt.Println("Daftar Produk:")
    for i, produk := range daftarProduk {
        fmt.Printf("%d. %s - Rp%.0f (Stok: %d)\n", i+1, produk.Nama, produk.Harga, produk.Stok)
    }

    // Loop untuk proses belanja
    for {
        fmt.Print("Masukkan nomor produk yang ingin dibeli: ")
        scanner.Scan()
        nomorProdukStr := scanner.Text()

        // Konversi nomor produk ke integer
        nomorProduk, err := strconv.Atoi(nomorProdukStr)
        if err != nil || nomorProduk < 1 || nomorProduk > len(daftarProduk) {
            fmt.Println("Nomor produk tidak valid!")
            continue
        }

        // Mengambil produk yang dipilih
        produkDipilih := daftarProduk[nomorProduk-1]

        // Memeriksa stok produk
        if produkDipilih.Stok <= 0 {
            fmt.Println("Stok produk habis!")
            continue
        }

        // Menanyakan jumlah yang ingin dibeli
        fmt.Print("Masukkan jumlah yang ingin dibeli: ")
        scanner.Scan()
        jumlahStr := scanner.Text()

        // Konversi jumlah ke integer
        jumlah, err := strconv.Atoi(jumlahStr)
        if err != nil || jumlah <= 0 {
            fmt.Println("Jumlah tidak valid!")
            continue
        }

        // Memeriksa apakah jumlah yang dibeli melebihi stok
        if jumlah > produkDipilih.Stok {
            fmt.Println("Jumlah melebihi stok!")
            continue
        }

        // Menghitung total harga
        hargaTotal := produkDipilih.Harga * float64(jumlah)

        // Menambahkan produk ke keranjang
        keranjangBelanja = append(keranjangBelanja, Produk{
            Nama:   produkDipilih.Nama,
            Harga:  produkDipilih.Harga,
            Stok:   jumlah,
        })

        // Mengurangi stok produk
        daftarProduk[nomorProduk-1].Stok -= jumlah

        // Menampilkan detail pembelian
        fmt.Printf("Produk %s (x%d) - Rp%.0f\n", produkDipilih.Nama, jumlah, hargaTotal)

        // Menghitung total belanja
        totalBelanja += hargaTotal

        // Menanyakan apakah ingin lanjut belanja
        fmt.Print("Lanjut belanja? (y/n): ")
        scanner.Scan()
        lanjutStr := scanner.Text()
        if strings.ToLower(lanjutStr) != "y" {
            break
        }
    }

    // Menampilkan total belanja
    fmt.Println("\nTotal belanja: Rp", totalBelanja)

    // Menampilkan detail keranjang belanja
    fmt.Println("\nKeranjang Belanja:")
    for _, produk := range keranjangBelanja {
        fmt.Printf("%s (x%d) - Rp%.0f\n", produk.Nama, produk.Stok, produk.Harga*float64(produk.Stok))
    }

    // Pesan terima kasih
    fmt.Println("\nTerima kasih telah berbelanja!")
}