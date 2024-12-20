package main

import "fmt"

func mengurutkan(a,b *int){
	if *a > *b{
		*a,*b = *b,*a
	}
}
func main(){
	var x,y int
	fmt.Scan(&x,&y)
	for x != y {
		mengurutkan(&x,&y)
		fmt.Print(x,y)
		break
	}
}