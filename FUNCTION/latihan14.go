package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	bilanganGanjil(n)

}
func bilanganGanjil(n int){
	var i int = 1
	ganjilSlave(n,i)
}
func ganjilSlave(n, i int){
	if !(i < n){
		if i % 2 != 0{
			fmt.Print(i, " ")
		}
	}else{
		if i % 2 != 0{
			fmt.Print(i, " ")
		}
	i++
	ganjilSlave(n, i)
	}
}