package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(sumFibonacci(n))
}
func sumFibonacci(n int)int{
	var f1,f2,f3,sum int
	f1 = 0
	f2 = 1
	sum = 0
	for i := 2; i <= n; i++{
		f3 = f2 + f1
		f1 = f2
		f2 = f3
		sum += f1
	}
	return sum
}