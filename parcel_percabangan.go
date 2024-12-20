package main

import "fmt"

func main(){
	var parcel, biaya, kilo, gram int

	//fmt.Scan(&parcel)
	//if parcel / 1000 && parcel % 1000 == 500 {
		//fmt.Println((parcel / 1000, "kg") + (parcel % 1000))
		//fmt.println((parcel / 1000 * 1000, "Rp") + (parcel % 1000 * 5))
	//}else{
		//fmt.Println((parcel / 1000,"kg") + (parcel % 1000))
		//fmt.Println((parcel / 1000 * 10000, "Rp") + (parcel % 1000 * 15))
	
	fmt.Scan(&parcel)
	gram = parcel % 1000

	if parcel < 10000 && gram >= 500 {
		kilo = parcel / 1000 * 10000
		biaya = kilo + (gram * 5)
		
	}else if parcel < 10000 && gram < 500 {
		kilo = parcel / 1000 * 10000
		biaya = kilo + gram * 15
	}else if parcel >= 10000 {
		biaya = parcel / 1000 * 10000
	}
	
	fmt.Println(biaya)
	
}