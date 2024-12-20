package main

import "fmt"

func main(){
	var jari, tinggi float64
	fmt.Scan(&jari,&tinggi)
	fmt.Printf("%.3f", 2 * luasLingkaran(jari,tinggi) + luasSelimut(jari,tinggi))
}
func luasLingkaran(r, t float64)float64{
	return 3.14 * (r * r)
}
func luasSelimut(r, t float64)float64{
	return 2 * 3.14 * r * t
}
/*PSEUDOCODE
program luas tabung
kamus 
	jari, tinggi : rill
algoritma
	input(jari,tinggi)
	output(2 * luasLingkaran(jari,tinggi + luasSelimut(jari,tinggi))
endprogram

function luasLingkaran(r,t : rill)-> rill
kamus
algoritma
	return 3.14 * r * r
endfunction

function luasSelimut(r,t : rill) -> rill
kamus
algoritma
	return 2 * 3.14 * r * t
endfunction
*/