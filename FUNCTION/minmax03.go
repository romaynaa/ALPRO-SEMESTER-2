//SEQUENTIAL SEARCH

package main

import "fmt"

const NMAX int = 10
type tabStr [NMAX]string

func main(){
	data := tabStr{"sepatu","bola","payung","kemeja"}
	var n int = 4
	var x string
	fmt.Scan(&x)
	fmt.Println(sequential_search(data,n,x))
	fmt.Println(sequential_search_idx(data,n,x))
}
func sequential_search(A tabStr, n int, x string)bool{
	var i int = 0
	var ketemu bool
	for i < n && !ketemu {
		ketemu = A[i] == x
		i++
	}
	return ketemu
}
func sequential_search_idx(A tabStr, n int, x string)int{
	var i int = 0
	var idx int = -1
	for i < n && idx == -1{
		if A[i] == x{
			idx = i
		}
		i++
	}
	return idx
}