package main

import "fmt"

func main(){
	var b, c, x float64
	
	fmt.Scan(&b, &c)
	if c > b {
		x = c - b
		fmt.Println("peningkatan sebesar", x)

	}else if c < b {
		x = b - c
		fmt.Println("penurunan sebesar", x)
	}else{
		fmt.Println("tetap")
	}

}