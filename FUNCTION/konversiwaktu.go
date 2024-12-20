package main

import "fmt"

func konversiWaktu(jam,menit,detik int)int{
	jam = jam * 3600
	menit = menit * 60
	detik = detik * 1
	return jam + menit + detik
}
func main(){
	var j,m,d int
	fmt.Scan(&j,&m,&d)
	fmt.Println(konversiWaktu(j,m,d))
}