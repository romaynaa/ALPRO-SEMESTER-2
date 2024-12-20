package main

import "fmt"

func main(){
	var gigi int
	var kopling, gas, menyala, jalan bool

	fmt.Scan(&gigi, &kopling, &gas)
	menyala = kopling || gas 
	jalan = gigi > 0 && kopling == false
	
	if menyala && jalan {
		fmt.Println("motor berjalan")
	}else if !menyala {
		fmt.Println("mesin mati")
	}else{
		fmt.Println("mesin menyala dan motor tidak berjalan")
	}
		
	
}