package main 

import "fmt"

const NMAX int = 100
type teman [NMAX]string

func main(){
	var t teman
	var n int
	fmt.Scan(&n)
	tulis(&t,n)
	cetak(t,n)
}
func tulis(A *teman, n int){
	for i := 0; i < n-1; i++{
		fmt.Scan(A[i])
	}
}
func cetak(A teman, n int){
	for i := 0; i < n-1; i++{
		fmt.Println(A[i])
	}
}