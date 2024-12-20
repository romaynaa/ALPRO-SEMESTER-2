package main

import "fmt"

func main(){
	var total_belanja, hasil int
	var kartu, diskon, cashback bool

	fmt.Scan(&total_belanja, &kartu)

	cashback = kartu && total_belanja >= 200000
	diskon = total_belanja >= 100000

	if diskon {
		hasil = total_belanja * 90 / 100
	}else if cashback {
		hasil = hasil - 75000
	}
	fmt.Println("kartu ?", kartu)
	fmt.Println("diskon ?", diskon)
	fmt.Println("cashback ?", cashback)
	fmt.Println("Rp", hasil)
}