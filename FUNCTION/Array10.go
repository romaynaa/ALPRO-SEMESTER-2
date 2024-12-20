package main

import "fmt"

const NMAX int = 10
type tabStr [NMAX]string

func main(){
	data := tabStr{"sepatu", "bola", "payung", "kemeja"}
	var n int = 4
	var x string
	fmt.Scan(&x)
	fmt.Println(sequentialSearch(data, n, x))

}
func sequentialSearch(A tabStr, n int, x string)bool{
	var i int = 0
	var ketemu bool
	for i < n && !ketemu{
		ketemu = A[i] == x
		i++
	}
	return ketemu
}