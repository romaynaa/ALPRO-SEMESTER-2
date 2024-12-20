package main

import "fmt"


func pecahUang(uang int, k10,k2,k1 *int){
	*k10 = uang / 10000
	*k2 = (uang - (*k10 * 10000)) /2000
	*k1 = (uang - (*k10 * 10000) - (*k2 * 2000))/1000
	
}
func main(){
	var k10000, k2000, k1000 int
	var money int
	fmt.Scan(&money)

	pecahUang(money, &k10000, &k2000, &k1000)
	fmt.Println(k10000," ", "lembar 10000", k2000, " ", "lembar 2000", k1000, " ", "lembar 1000")
	
}