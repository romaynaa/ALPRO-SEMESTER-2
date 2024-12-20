package main

import "fmt"

func main(){
	var n int
	var status, kartu, diskon, cashback bool
	fmt.Scan(&n, &status)

	if n >= 100000 && status {
		fmt.print(kartu)
	}else if n >= 100000 {
		fmt.Print(diskon)
	}else if n >= 200000 && kartu {
		fmt.Print(cashback)
	}
}