package main

import "fmt"

func reamur(c float64) float64{
	r := c * 0.80
	return r
}
func fahrenheit(c float64) float64{
	f := (9*c) / 5 + 32
	return f
}
func kelvin(c float64) float64{
	k := c + 273
	return k
}

func main(){
	var c_awal, c_akhir, step float64
	fmt.Scan(&c_awal, &c_akhir, &step)
	
	fmt.Printf("%10s %10s %10s %10s", "Celcius","Reamur", "Fahrenheit", "Kelvin")
	fmt.Println()
	for c_awal <= c_akhir {
		C := c_awal
		R := reamur(c_awal)
		F := fahrenheit(c_awal)
		K := kelvin(c_awal)
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", C, R, F, K)
		c_awal = c_awal + step
	}
}
