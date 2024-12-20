package main

import "fmt"

func main(){
	var n int
	var status, kartu, diskon, cashback bool
	fmt.Scan(&n, &status)

	if diskon {
		diskon = (n >= 100000)
	}else if kartu {
		kartu = status && diskon
	}else if cashback {
		cashback = (n >= 200000) && kartu
	
	}
	fmt.Println(kartu, diskon, cashback)
}