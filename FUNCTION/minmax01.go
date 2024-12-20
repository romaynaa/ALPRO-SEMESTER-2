package main

import "fmt"

const NMAX int = 2001
type tabInt [NMAX]int

func main(){
	data := tabInt{4,0,2,5,7}
	var nData int
	nData = 5
	fmt.Println(maximum(data,nData))
}
func maximum(d tabInt, n int)int{
	var max int
	max = d[0]
	for i := 1; i < n; i++{
		if d[i] > max{
			max = d[i]
		}
	}
	return max
}
func cetak(d tabInt, n int){
	for i := 0; i < n; i++{
		fmt.Printf("%d\n", d[i])
	}
	fmt.Println()
}