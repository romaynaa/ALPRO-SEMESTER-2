package main

import "fmt"

func hitungLuasKelilingLingkaran(r float64, l,k *float64){
	*l = 3.14 * r * r
	*k = 2 * 3.14 * r
}
func hitungLuasKelilingPersegi(s float64, l,k *float64){
	*l = s * s
	*k = 4 * s
}
func hitungTotal(IL,IP,kL,kP float64, totLuas,totKel *float64){
	*totLuas = IL +IP
	*totKel = kL + kP
}
func main(){
	var radius, sisi, luasL, kelL, luasP, kelP, totalL, totalK float64
	fmt.Scan(&radius, &sisi)
	if radius != 0 && sisi != 0 {
		fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
	}
	for radius != 0 && sisi != 0 {
		hitungLuasKelilingLingkaran(radius, &luasL, &kelL)
		hitungLuasKelilingPersegi(sisi, &luasP, &kelP)
		hitungTotal(luasL,luasP,kelL,kelP,&totalL,&totalK)
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n",radius, sisi, luasL, kelL, luasP, kelP, totalL, totalK)
		fmt.Scan(&radius,&sisi)
	}
}