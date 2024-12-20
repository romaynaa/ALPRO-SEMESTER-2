package main

import "fmt"

func main(){
	var tahun int
	fmt.Scan(&tahun)

	if tahun % 400 == 0 || tahun % 4 == 0{
		fmt.Println(tahun, "adalah tahun kabisat")
	}else{
		fmt.Println(tahun, "bukanlah tahun kabisat")
	}
}