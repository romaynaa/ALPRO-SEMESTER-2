package main

import "fmt"

func main(){
	
	var n, bil, hasil int

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
	fmt.Scan(&bil)
	hasil += bil
	}
	fmt.Println(hasil)
}