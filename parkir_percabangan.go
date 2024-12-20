package main

import "fmt"

func main(){
	var h1, m1, h2, m2 int
	fmt.Scan(&h1, &m1, &h2, &m2)
	if h2 < h1 {
		fmt.Print(h2 + 12 - h1, "jam")
	}else{
		fmt.Print(h2 - h1, "jam")
	}
	fmt.Print(m2 - m1, "menit")
	
}