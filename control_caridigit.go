package main

import "fmt"

func main(){

	var x, n , digit int
	var hasil bool

	fmt.Scan(&x, &n)

	for n > 0 {
		digit = n % 10
		if digit == x {
			hasil = true
		}
		n = n / 10
	}
	fmt.Println(hasil)
}
//pseudocode
// program cari digit
// kamus : x, n, digit integer
// 		hasil boolean
// algoritma : 
// 	input (x,n)
// 	while n > 0 do
// 	digit <- n mod 10
// 		if digit == x then 
// 		hasil == true
// 	endif
// 	n <- n div 10
// 	endwhile
// 	output(hasil)
// endprogrM
