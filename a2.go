package main

import "fmt"

func main(){
	var member string
	var a, b, c, d, e int
	var c1, c2, c3, c4, c5 bool
	var cash, diskon float64

	fmt.Scan(&member, &a, &b, &c, &d, &e)
	c1 = a % 2 == 0
	c2 = b % 2 == 0
	c3 = c % 2 == 0
	c4 = d % 2 == 0
	c5 = e % 2 == 0

	if c1 == true && c2 == true && c3 == true && c4 == true && c5 == true {
		cash = 3.1 * (a+b+c)
		fmt.Println (cash)
	}
}
