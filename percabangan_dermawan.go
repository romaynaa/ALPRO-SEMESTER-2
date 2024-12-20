package main

import "fmt"

func main(){
	var n_minggu, n1, n2, n3, n4, n5, uang int

	fmt.Scan(&n_minggu)
	if n_minggu == 4 {
		fmt.Scan(&n1, &n2, &n3, &n4)
		uang = n1 + n2 + n3 + n4
		fmt.Println(uang)
	}else if n_minggu == 5 {
		fmt.Scan(&n1, &n2, &n3, &n4, &n5)
		uang = n1 + n2 + n3 + n4 + n5
		fmt.Println(uang)
	}
}