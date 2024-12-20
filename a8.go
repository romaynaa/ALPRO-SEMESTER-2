package main

import "fmt"

func main(){
	var r int
	
	fmt.Scan(&r)
	if r % 7 == 0 {
		fmt.Println(22 * r / 7 * r)
	}else{
		fmt.Println(float64(r * r) * 3.14)
	}
	
}