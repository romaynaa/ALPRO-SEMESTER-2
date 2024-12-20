package main

import "fmt"

func main(){
	var b1, b2, b3, b4, b5 int
	var status bool

	fmt.Scan(&b1, &b2, &b3, &b4, &b5)
	status = (b1 < b2) && (b2 < b3) && (b3 < b4) && (b4 < b5)
	
	fmt.Println(status)
}