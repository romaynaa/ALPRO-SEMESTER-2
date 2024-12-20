package main 

import "fmt"

const NMAX int = 100
type pemain struct{
	tim1, tim2 int
}
type tabGol [NMAX] pemain
func main(){
	var t tabGol
	var n int
}
func bacaData(t *tabgol, n *int){
	fmt.Scan(&n)
	for i := 0; i < *n; i++{
		fmt.Scan(&t.pemain.tim1, &t.pemain.t2)
		if t.pemain.tim1 < 0 && t.pemain.tim2 < 0{
			break
		}
	}
}
func rataan(t tabGol, n int)float64{
	return float64(bacaData())
}