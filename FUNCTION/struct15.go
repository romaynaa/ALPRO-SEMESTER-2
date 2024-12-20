package main 

import "fmt"

func main(){
	type buku struct{
		judul, nama_penulis string
		tahun_terbit, harga int
	}
	var b buku
	
	fmt.Scan(&b.judul, &b.nama_penulis, &b.tahun_terbit, &b.harga)
	fmt.Println(b.judul, b.nama_penulis, b.tahun_terbit, b.harga)
	// bil = 1
	// b.judul = "algoritma keren"
	// b.nama_penulis = "Suo"
	// b.tahun_terbit = 2023

	// fmt.Println(bil)
	// fmt.Println(b.judul,b.nama_penulis,b.tahun_terbit
	
}