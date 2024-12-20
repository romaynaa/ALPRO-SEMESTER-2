package main

import "fmt"

func main(){
	var n, x, y, i int
	fmt. Scan(&x, &y)
	i = 0
	n = x
	for n >= y {
		n = n - y
		i = i + 1

	}
fmt.Println(x, "mod", y, "=", n)
fmt.Println(x, "div", y, "=", i)
}