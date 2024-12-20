package main

import "fmt"

func main(){
	var x, y, a, f float64

	fmt.Scan(&x, &y, &a)
	f = (2*x*(1-a)/3*y*y) + 5*y/4*x*(1-a)
	fmt.Println(f)
}