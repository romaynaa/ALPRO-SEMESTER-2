package main 

import "fmt"

func main(){
	var jam, menit, detik int
	fmt.Scan(&jam,&menit,&detik)
	fmt.Println(konversiWaktu(jam,menit,detik))

}
func konversiWaktu(j,m,d int)int{
	j = j * 3600
	m = m * 60
	d = d * 1
	return j + m + d
}