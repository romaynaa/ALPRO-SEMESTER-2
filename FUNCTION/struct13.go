package main

import (
	"fmt"
	"math"
)
type titik struct{
	x,y float64
}
func main(){
	var t1,t2 titik
	fmt.Scan(&t1.x,&t1.y,&t2.x,&t2.y)
	fmt.Printf("%.2f\n",jarak(t1,t2))
	
}
func jarak(p1,p2 titik)float64{
	return akar((p1.x-p2.x)*(p1.x-p2.x)+(p1.y-p2.y)*(p1.y-p2.y))
}
func akar(a float64)float64{
	return math.Sqrt(a)
}