package main

import "fmt"

func main(){

	var total_belanja float64
	var bersedia, Kartu, Diskon, Cashback bool
	Kartu, Diskon, Cashback = true, true, true

	fmt.Scan(&total_belanja, &bersedia)

	if total_belanja < 100000 {
		Diskon = false
	} else {
		total_belanja -= total_belanja * 0.1
		
		if bersedia {
			if total_belanja >= 200000 {
				total_belanja -= 75000
			} else {
				Cashback = false
			}
		} else {
			Cashback = false
			Kartu = false
		}
	}

	fmt.Println("Kartu? ", Kartu)
	fmt.Println("Diskon? ", Diskon)
	fmt.Println("Cashback? ", Cashback)
	fmt.Println("Rp. ", int(total_belanja))
}