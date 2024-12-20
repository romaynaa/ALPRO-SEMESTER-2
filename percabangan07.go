package main

import "fmt"

func main(){
	var tot_belanja, hasil float64
	var kupon bool
	fmt.Scan(&tot_belanja, &kupon)
	
	//if kupon == true && tot_belanja >= 100000 {
		//hasil = tot_belanja - ((tot_belanja * 0.05) - tot_belanja) - (float64(tot_belanja / 100000) * 10000)
	//}else if kupon == false && tot_belanja < 100000{
		//hasil = tot_belanja
	//}else if kupon == false && tot_belanja >= 100000{
		//hasil = tot_belanja + (float64(tot_belanja / 100000) * 10000) 
	//}else if kupon == true && tot_belanja < 100000{
		//hasil = tot_belanja - (tot_belanja * 0.05)
	//}
	//fmt.Println(hasil)
	if kupon {
		hasil = tot_belanja - (tot_belanja * 0.05)
	}else{
		hasil = tot_belanja
	}
	if tot_belanja >= 100000 {
		hasil = hasil - (tot_belanja / 100000) * 10000 
	}
	fmt.Println(hasil)
}