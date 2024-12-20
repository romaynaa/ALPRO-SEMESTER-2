package main

import (
	"fmt"
	"math"
)

func panjangX(R, T float64) float64{
	return R * math.Cos(T)
}
func panjangY(R,T float64) float64{
	return R * math.Sin(T)
}

func main(){
	var R, T, pX, pY float64
	fmt.Scan(&R, &T)

	T = T * math.Pi/180.0
	pX = panjangX(R,T)
	pY = panjangY(R,T)

	fmt.Printf("%.2f %.2f\n", pX, pY)
}