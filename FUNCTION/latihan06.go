package main

import "fmt"

func main(){
	var t_awal int
	var member bool
	fmt.Scan(&t_awal,&member)
	fmt.Println(hitungDiskon(member, t_awal))
}
func hitungDiskon(m bool, t int)int{
	var diskon int
	if t > 100000 && m == true{
		diskon = t * 10 / 100
	}else if t > 100000 && m == false{
		diskon = t * 5 / 100
	}else{
		diskon = 0
	}
	return t - diskon
}