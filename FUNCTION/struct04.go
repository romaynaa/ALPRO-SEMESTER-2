package main

import (
	"fmt"
	"math"
)

type titik struct{
	x, y float64
}
type lingkaran struct{
	titik titik
	radius, luas, keliling float64
}
func main(){
	var c1, c2 lingkaran
	var x,y,radius float64
	fmt.Scan(&x,&y,&radius)
	//membuat lingkaran baru c1
	c1 = lingkaran_baru(x,y,radius)
	fmt.Scan(&x,&y,&radius)
	//membuat lingkaran baru c2
	c2 = lingkaran_baru(x,y,radius)

	//menghitung luas c1
	c1 = hitung_luas(&c1)
	//menghitung keliling c1
	c1 = hitung_keliling(&c1)

	c2 = hitung_luas(&c2)
	c2 = hitung_keliling(&c2)

	cetak_data(c1)
	cetak_data(c2)

}
func lingkaran_baru(x,y,r float64)lingkaran{
	var l lingkaran
	var tp titik
	l.tp.x = x
	l.tp.y = y
	l.radius = r
	return l
	
}
func hitung_luas(l *lingkaran){
	const pi float64 = 3.14
	l.luas = math.Pi * l.radius *l.radius
}
func hitung_keliling(l *lingkaran){
	const pi float64 = 3.14
	l.keliling = 2 * math.Pi * l.radius
}
func cetak_data(l lingkaran){
	fmt.Printf("lingkaran berpusat di titik (%v,%v) dan beradius %v memiliki luas %.2f dan keliling sebesar %.2f ", l.titik.x)
}