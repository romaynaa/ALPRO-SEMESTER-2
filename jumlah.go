package main 

import "fmt"

func main(){
	var n, bil, digit, hasil int
	fmt.Scan(&n, &bil)

	for i := 1; i <= n; i++ {
		digit = bil % 10
		hasil = digit + hasil
		bil = bil / 10
		
	}
	fmt.Println(hasil)
}