package main

import "fmt"

func main(){
	var x, n, digit int
	var status bool
	fmt.Scan(&x, &n)
	status = false
	for n > 0 {
		digit = n % 10 
		if digit == x {
			status = true
		}
		n = n / 10
	}
	fmt.Println(status)
}