package main

import "fmt"

func temperatur(celcius float64)float64{
	return (9/5 * celcius) + 32
}
func main(){
	var n int
	var c float64
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&c)
		fmt.Println(c," ","Celcius"," ", temperatur(c)," ","Fahrenheit")
	}
}