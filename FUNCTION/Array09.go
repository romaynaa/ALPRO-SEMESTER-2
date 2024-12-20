package main 

import "fmt"

const NMAX int = 2021

type tabInt[NMAX] int

func main(){
	data := tabInt{1,0,2,5,7}
	var ndata int = 5
	fmt.Println(maximum(data, ndata))
	fmt.Println(idx_max(data, ndata))
}
func idx_max(A tabInt, n int)int{
	var i, idx_max int
	idx_max = 0
	for i = 1; i < n; i++{
		if A[i] > A[idx_max]{
			idx_max = i
		}
	}
	return idx_max
}
func maximum(A tabInt, n int)int{
	var i, max int
	max = A[0]
	for i = 1; i < n; i++{
		if A[i] > max{
			max = A[i]
		}
	}
	return max
}
func cetakData(A tabInt, n int){
	for i := 0; i < n; i++{
		fmt.Printf("%d", A[i])
	}
	fmt.Println()
}