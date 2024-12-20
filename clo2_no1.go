package main

import "fmt"

func main(){
	var n, voucher, hasil, digit1 int
	fmt.Scan(&n, &voucher)

	for i := 1; i <= n; i++ {
	digit1 = (voucher / 10000)
	hasil = digit1 + hasil
	voucher = voucher 
	}
	fmt.Println(hasil * 20000)
}
