package main

import "fmt"

func main(){
	var r, luas1 int
	var luas2 float64
	fmt.Scan(&r)
	if r % 7 == 0 {
		luas1 = (22 * r / 7 * r)
		fmt.Println(luas1)
	}else{
		luas2 = float64(r * r) * 3.14
		fmt.Println(luas2)
	}
	
}