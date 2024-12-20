//PROGRAM PERSEGI

package main

import "fmt"

type rectangle struct{
	lenght, widht int
	color string
	property geometry
}
type geometry struct{
	luas, keliling int
}
func main(){
	var persegi rectangle
	isiData(&persegi)
	hitung(&persegi)
	fmt.Print(persegi.property.luas, " ")
	fmt.Print(persegi.property.keliling)

}
func isiData(persegi *rectangle){
	fmt.Scan(&persegi.lenght,&persegi.widht,&persegi.color) 
}
func hitung(persegi *rectangle){
	persegi.property.luas = persegi.lenght * persegi.widht
	persegi.property.keliling = 2 * (persegi.lenght + persegi.widht)
}