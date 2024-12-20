package main

import "fmt"

func cetak_rerata(n, jumlah,i int, rata float64){
	if i > n{
		return 
	}
	jumlah += i
	rata = float64(jumlah)/float64(i)
	fmt.Println(jumlah, rata)
	cetak_rerata(n, jumlah, i+1, rata)
}
func main(){
	var n int
	fmt.Scan(&n)
	cetak_rerata(n,0,1,0)
}