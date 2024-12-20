package main

import "fmt"

const NMAX int = 3
type tabInt struct{
	info[NMAX] int
	ndata int
}
func main(){
	var d tabInt
	bacaData(&d)
	fmt.Printf("Jumlah bilangan %d\n", hitungJumlah(d))
	fmt.Printf("Rata-rata bilangan %.2f\n", hitungRerata(d))
}
func bacaData(d *tabInt) {
	var i int = 0
	fmt.Scan(&d.ndata)
	if d.ndata > NMAX{
		d.ndata = NMAX
	}
	for i < d.ndata{
		fmt.Scan(&d.info[i])
		i++
	}
}
func hitungJumlah(d tabInt)int{
	var jumlah int = 0
	for i := 0; i < d.ndata; i++{
		jumlah += d.info[i]
	}
	return jumlah
}
func hitungRerata(d tabInt)float64{
	return float64(hitungJumlah(d))/ float64(d.ndata)
}