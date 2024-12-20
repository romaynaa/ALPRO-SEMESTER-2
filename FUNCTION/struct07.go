package main

import "fmt"

type balok struct{
	panjang, lebar, tinggi int
	volume, luas int
}
func main(){
	var t balok
	fmt.Scan(&t.panjang,&t.lebar,&t.tinggi)
	luasBalok(&t)
	volumeBalok(&t)
	fmt.Println(t.luas,t.volume)
}
func luasBalok(t *balok){
	t.luas = 2 * (t.panjang * t.lebar + t.panjang * t.tinggi + t.lebar * t.tinggi)
}
func volumeBalok(t *balok){
	t.volume = t.panjang * t.lebar * t.tinggi
}