package main 

import "fmt"

type rectangle struct{
	lenght, widht int
	color string
	property geometry
}
type geometry struct{
	area, perimeter int
}
func main(){
	var p rectangle
	isiData(&p)
	hitung(&p)
	fmt.Print(p.property.area, " ", p.property.perimeter)
}
func isiData(p *rectangle){
	fmt.Scan(&p.lenght, &p.widht, &p.color)
}
func hitung(p *rectangle){
	p.property.area = p.lenght * p.widht
	p.property.perimeter = 2 * (p.lenght + p.widht)
}