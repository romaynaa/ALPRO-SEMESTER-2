package main

import "fmt"

func barisan(n int)int{
	if n <= 3{
		return 1
	}else{
		return barisan(n-1) + barisan(n-2) + barisan(n-3)
	}
}
func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(barisan(n))
}