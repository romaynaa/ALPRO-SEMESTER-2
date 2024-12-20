package main

import "fmt"

func main(){
	var n int
	var hariDanAcara string
	var kosong bool

	fmt.Scan(&n)
	kosong = true 
	if n <= 0 {
		fmt.Println("Angka Invalid")
	}else{
		if n % 7 == 1 {
			hariDanAcara = "senin"
		}else if n % 7 == 2 {
			hariDanAcara = "selasa"
		}else if n % 7 == 3 {
			hariDanAcara = "rabu rapat rutin"
			kosong = false
		}else if n % 7 == 4 {
			hariDanAcara = "kamis"
		}else if n % 7 == 5 {
			hariDanAcara = "jumat"
		}else if n % 7 == 6 {
			hariDanAcara = "sabtu"
		}else if n % 7 == 0 {
			hariDanAcara = "minggu"
		}
		if n % 3 == 0 {
			hariDanAcara = hariDanAcara + "konsultasi kesehatan"
			kosong = false
		}
			if n % 4 == 0 {
			hariDanAcara = hariDanAcara + "mengecek keadaan orangtua"
			kosong = false
		}
			if kosong {
			hariDanAcara = hariDanAcara + "kosong"
		}
fmt.Println(hariDanAcara)
}
}


















































