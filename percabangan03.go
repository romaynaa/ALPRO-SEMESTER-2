package main

import "fmt"

func main(){
	var modal, jual int
	fmt.Scan(&modal,&jual)

	if modal > jual {
		fmt.Println("untung",modal - jual)
	}else if jual > modal{
		fmt.Println("rugi", jual - modal )
	}else{
		fmt.Println("tidak rugi dan tidak untung")
	}
}