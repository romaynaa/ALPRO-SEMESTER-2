package main

import "fmt"

type titik struct{
	x,y float64
}
type lingkaran struct{
	titik titik
	radius, luas, keliling float64
}
func main(){
	var c1, c2 lingkaran
	var x,y,radius float64
	fmt.Scan(&x,&y,&radius)
	c1 = ----

	fmt.Scan(&x,&y,&radius)
	c2 = ---

	//hitung luas c1

	//hitung luas c2

	//hitung keliling c1

	//hitung keliling c2

	//cetak data
	cetak_data(c1)
	cetak_data(c2)
}
func lingkaran_baru(x,y,r float64)lingkaran{
	
}
