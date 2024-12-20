package main 

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	faktorBilangan(n)


}
func faktorBilangan(n int){
	var i int = 1
	faktorBilanganSlave(n, i)
}
func faktorBilanganSlave(n, i int){
	if !(i <= n){
		if n % i == 0{
			fmt.Print(i, " ")
		}
	}else{
		if n % i == 0{
			fmt.Print(i, " ")
		}
	i++
	faktorBilanganSlave(n,i)
	}
}