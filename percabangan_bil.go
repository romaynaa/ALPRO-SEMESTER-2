package main

import "fmt"

func main(){
	var a, b, c, tampung int
	fmt.Scan(&a, &b, &c)

	if b > c {
		tampung = b
		b = c
		c = tampung
	}
	if a > c {
		tampung = a
		a = c
		c = tampung
	}
	if a > b {
		tampung = a
		a = b
		b = tampung
	
	}
	fmt.Println(a, b, c)
}