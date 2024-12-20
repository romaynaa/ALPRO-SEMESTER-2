package main 

import "fmt"

const NMAX int = 2001
type tabInt [NMAX]int

func main(){
	data := tabInt{4,0,2,5,7}
	var nData int
	nData = 5
	fmt.Println(idx_max(data,nData))

}
func idx_max(d tabInt, n int)int{
	var idx int
	idx = 0
	for i := 1; i < n; i++{
		if d[i] > d[idx] {
			idx = i
		}
	}
	return idx
}
func cetak(d tabInt, n int){
	for i := 0; i < n; i++{
		fmt.Printf("%d", d[i])
	}
	fmt.Println()

















}