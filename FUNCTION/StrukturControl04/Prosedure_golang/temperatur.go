package main

import "fmt"

func konversikanCelcius(c float64, r,f,k *float64){
	*r = c * 0.80
	*f = (9 * c) / 5 + 32
	*k = c + 273.15
}
func main(){
	var celcius, reamur, fahrenheit, kelvin float64
	fmt.Scan(&celcius)
	konversikanCelcius(celcius, &reamur, &fahrenheit, &kelvin)
	fmt.Print(reamur, "R", " ", fahrenheit, "F", " ", kelvin, "K")
}