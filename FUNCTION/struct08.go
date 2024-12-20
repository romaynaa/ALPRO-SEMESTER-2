//PROGRAM JARAK TITIK

package main

import (
	"fmt"
	"math"
)

type titik struct{
	x,y float64
}
func main(){
	var P1,P2 titik
	fmt.Scan(&P1.x, &P1.y, &P2.x, &P2.y)
	fmt.Printf("%.2f",jarak(P1,P2))

}
func jarak(p1,p2 titik)float64{
	var jarak float64
	jarak = akar(math.Pow(p1.x - p2.x, 2) + math.Pow(p1.y - p2.y, 2))
	return jarak
}
func akar(x float64)float64{
	return math.Sqrt(x)
}