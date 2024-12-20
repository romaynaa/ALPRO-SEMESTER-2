package main 

import "fmt"

func tabung(r, t, hasil float64)float64{
	hasil = 3.14 * (r * r) * t
	return hasil
}
func main(){
	var jarijari, tinggi, hasil float64
	fmt.Scan(&jarijari,&tinggi)
	fmt.Println(tabung(jarijari,tinggi,hasil))
}