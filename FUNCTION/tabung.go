package main

import "fmt"

func volumeTabung(jarijari, tinggi float64)float64 {
	hasil := 3.14 * (jarijari * jarijari) * tinggi
	return hasil
}

func main(){
	var r, t, volume float64
	fmt.Scan(&r, &t)

	volume = volumeTabung(r,t)
	fmt.Println(volume)
}