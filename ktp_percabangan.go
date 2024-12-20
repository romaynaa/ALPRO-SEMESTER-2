package main

import "fmt"

func main(){
	var usia int
	var membuat_kk bool
	
	fmt.Scan(&usia, &membuat_kk)
	if usia >= 17 && membuat_kk  {
		fmt.Println("bisa membuat ktp")
	}else{
		fmt.Println("belum bisa membuaat ktp")
	}
}