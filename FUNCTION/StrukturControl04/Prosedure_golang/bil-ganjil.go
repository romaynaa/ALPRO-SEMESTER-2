package main

import "fmt"

func bilanganGanjil(n int)int{
	if n <= 1 {
		return 1
	}
	if n % 2 != 0 {
		fmt.Println((n))
	}
	return bilanganGanjil(n - 1)
}
func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(bilanganGanjil(n))
}