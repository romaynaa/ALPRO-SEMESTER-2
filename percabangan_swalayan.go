package main

import "fmt"

func main(){
	//var member string
	//var a, b, c, d, e int
	//var cashback, diskon float64
	//fmt.Scan(&member, &a, &b, &c, &d, &e)
	//if member == "yes" {
	//	a%2 == 1 && b%2 == 1 && c%2 == 1 && d%2 == 1 && e%2 == 1
		//fmt.Println(cashback == 3.1 * (a+b+c))
	//}
		var member string
		var a, b, c, d, e int
		var diskon, cash float64
		
		fmt.Scan(&member, &a, &b, &c, &d, &e)
		if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
			cash = 3.1 * float64 (a+b+c)
		}else if a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0 {
			diskon = 1.7 * float64 (c+d+e)
		}else{
			cash = 3.1 * float64 (a+b+c)
			diskon = 1.7 * float64 (c+d+e)
		}
		if member == "yes" {
			cash = cash + (cash * 0.15)
			diskon = diskon + (diskon * 0.15)
		}
		if cash > 35 {
			cash = 35
		}
		if diskon > 50 {
			diskon = 50
		}

	fmt.Println(cash,"%", diskon,"%")
}