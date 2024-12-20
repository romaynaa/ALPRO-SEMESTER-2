package main

import "fmt"

func reamur(C float64)float64{
	Reamur := C * 0.80
	return Reamur
}
func fahrenheit(C float64)float64{
	Fahrenheit := C * 1.8 + 32
	return Fahrenheit
}
func kelvin(C float64)float64{
	Kelvin := C + 273
	return Kelvin
}
func main(){
	
	var Cawal, Cakhir, step float64
	fmt.Scan(&Cawal, &Cakhir, &step)
	fmt.Printf("%10s %10s %10s %10s", "Celcius", "Reamur", "Fahrenheit", "Kelvin")
	fmt.Println()
	for Cawal <= Cakhir{
		C := Cawal
		R := reamur(Cawal)
		F := fahrenheit(Cawal)
		K := kelvin(Cawal)
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", C, R, F,K)
		Cawal = Cawal + step
	}
}