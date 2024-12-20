package main

import "fmt"

func main(){
	var n, m, i int
	fmt.Scan(&n,&m)
	for i = n; ; i++ {
		if i % n == 0 && i % m == 0 {
			fmt.Print(i)
			break
		}
	}
}