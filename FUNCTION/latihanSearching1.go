package main

import "fmt"

const NMAX int = 12
type tabInt [NMAX]int

func main(){
	var data tabInt
	var x, ndata int
	fmt.Scan(&x)
	ndata = 0
	for x != -5313541 && ndata < NMAX{
		if x == 0{
			sorting(&data, ndata)
			fmt.Println(median(data, ndata))
		}else{
			data[ndata] = x
			ndata++
		}
		fmt.Scan(&x)
	}
}
func sorting(A *tabInt, n int){
	var i, idx_min, j, temp int
	for i = 1; i <= n-1; i++ {
		idx_min = i - 1
		for j = i; j <= n-1; j++{
			if A[idx_min] > A[j]{
				idx_min = j
			}
		}
		temp = A[idx_min]
		A[idx_min] = A[i-1]
		A[i-1] = temp 
	}
	/*
	var i, idx_min, j, temp int
	for i = 1; i <= n-1; i++{
		idx_min = i - 1
		for j = i; j <= n-1; j++{
			if A[idx_min] > A[j]
			idx_min = j
		}
	}
		temp = A[idx_min]
		A[idx_min] = A[i-1]
		A[i-1] = temp
	*/
}
func median(A tabInt, n int) float64{
	var mid int = n/2
	if n % 2 == 0{
		return float64(A[mid-1] + A[mid]) / 2.0
	}else{
		return float64(A[mid])
	}
}