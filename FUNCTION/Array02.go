package main

import "fmt"

const NMAX int = 49
type teman[NMAX] string

func main(){
	var t teman
	var n int
	fmt.Scan(&n)
	bacaArray(&t,n)
	cetakArray(t,n)


}
func bacaArray(nama *teman, m int){
	for i := 0; i < (m-1); i++{
		fmt.Scan(nama[i])
	}
}
func cetakArray(nama teman, m int){
	for i := 0; i < (m-1); i++{
		fmt.Println(nama[i])
	}
}