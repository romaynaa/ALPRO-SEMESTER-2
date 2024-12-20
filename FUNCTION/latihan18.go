package main

import "fmt"

type balok struct{
	panjang, lebar, tinggi int
	volume, luas int
}
func main(){
	var bb balok
	fmt.Scan(&bb.panjang,&bb.lebar,&bb.tinggi)
	hitungLuas(&bb)
	hitungVolume(&bb)
	fmt.Println(bb.luas, bb.volume)
}
func hitungLuas(x *balok){
	x.luas = 2 * (x.panjang*x.lebar* + x.panjang*x.tinggi + x.lebar*x.tinggi)
}
func hitungVolume(x *balok){
	x.volume = x.panjang * x.lebar * x.tinggi
}