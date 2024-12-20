package main

import "fmt"

func main(){
	var jari, tinggi float64
	fmt.Scan(&jari,&tinggi)
	fmt.Printf("%.2f\n",tabung(jari,tinggi))

}
func tabung(r,t float64)float64{
	return 3.14 * (r * r) * t
}
/*PSEDEUCODE
program tabung
kamus
	jari, tinggi : real
algoritma
	input(jari,tinggi)
	output(tabung(jari,tinggi))
endprogram

function tabung(r,t : real) -> real
kamus 
algoritma
	return 3.14 * (r*r) * t
endfunction
*/