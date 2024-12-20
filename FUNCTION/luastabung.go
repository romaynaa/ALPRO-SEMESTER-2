package main

import "fmt"

func luasLingkaran(r float64)float64{
	return 3.14 * (r * r)
}
func luasSelimut(r,t float64)float64{
	return 2 * 3.14 * r *t
}
func main(){
	var jarijari, tinggi float64
	fmt.Scan(&jarijari, &tinggi)
	fmt.Println(2 * luasLingkaran(jarijari) + luasSelimut(jarijari,tinggi))
}