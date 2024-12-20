package main

import "fmt"

func fibo(n int)int{
	var i,u0,u1,un int
	u0 = 0
	u1 = 1
	
	for i = 1;i <= n; i++{
		un = u0 + u1
		u1 = u0
		u0 = un 
	}
	return i
}
func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(fibo(n))
}
