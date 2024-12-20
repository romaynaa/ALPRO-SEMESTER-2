package main 

import "fmt"

func main(){
	type bil_bul int
	var bil1, bil2, bil3 bil_bul
	var bil4 int

	fmt.Scan(&bil1,&bil2,&bil3)
	bil4 = int(bil1) + int(bil2) + int(bil3)
	fmt.Println(bil4)
}