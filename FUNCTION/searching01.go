package main

import "fmt"

const NMAX int = 20
type tabInt [NMAX]int

func main(){
	var data tabInt
	var nData int
	baca_data(&data, &nData)
	fmt.Println("Suara masuk:", jumlah(data, nData))
}
func baca_data(A *tabInt, n *int){
	var i int = 0
	for i < *n && i < NMAX{
		fmt.Scan(&A[i])
	}
	*n = i
}
func jumlah(A tabInt, n int)int{
	var jumlah int = 0
	for i := 0; i < n; i++{
		jumlah += A[i]
	}
	return jumlah
}