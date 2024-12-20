package main

import "fmt"

func main(){
	var n, digit, hasil int
	fmt.Scan(&n)

	for n > 0 {
		
		digit = n % 10 
		if digit > hasil {
			hasil = digit
		}
		n = n / 10
	}
	fmt.Println(hasil)
}