package main

import "fmt"

func main(){
	var n int
	var status, kartu, diskon, cashback bool
	fmt.Scan(&n, &status)

	if n > 100000 {
		fmt.Print(diskon)
	}else if status && diskon {
		fmt.Print(kartu)
	}else if (n > 200000) && kartu {
		fmt.Print(cashback)
	}
	fmt.Println(kartu, diskon, cashback)
}}else if kartu && cashback == true || diskon == true  {
	tot_belanja = n - (n * (10 / 100)) - 75000