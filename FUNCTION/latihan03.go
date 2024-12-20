package main

import "fmt"

func main(){
	var n, i int
	var celcius float64
	fmt.Scan(&n)
	i = 1
	for i <= n{
		fmt.Scan(&celcius)
		fmt.Print(celcius," ", "Celcius"," ", "="," ", temperatur(celcius)," ", "Fahrenheit")
	i++
	}
	
}
func temperatur(c float64)float64{
	return 1.8 * c + 32
}