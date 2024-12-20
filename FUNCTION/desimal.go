package main

import "fmt"

func main(){
	var desimal, sisa int
	var biner string
	fmt.Scan(&desimal)
	for desimal > 0{
		if desimal % 2 != 0{
			sisa = 1
		}else{
			sisa = 0
		}

	}
	fmt.Println(biner)
}