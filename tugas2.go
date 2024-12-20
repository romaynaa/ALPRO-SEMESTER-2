package main

import "fmt"

func main(){

	var bil, hasil int

	fmt.Scan(&bil)
	hasil = 0
	for bil != -1 {
	hasil = hasil + bil
	fmt.Scan(&bil)
		
	}
	fmt.Println(hasil)
}