package main

import "fmt"

func main(){
	var n int
	var biner string
	fmt.Scan(&n)
	for n > 0 {
		if n % 2 == 0 {
			biner = "0" + biner
		}else if n % 2 == 1 {
			biner = "1" + biner
		}
		n = n / 2
	}
	print(biner)
}