package main 

import "fmt"

func main(){
	var n, a, b int
	var hasil float64
	fmt.Scan(&n)
	for i := 1; i <= n; i++{
		fmt.Scan(&a,&b)
		luasSegitiga(a,b,&hasil)
		fmt.Println(hasil)
	}
}
func luasSegitiga(x,y int, hasil *float64){
	*hasil = float64(x) * float64(y) / 2
}