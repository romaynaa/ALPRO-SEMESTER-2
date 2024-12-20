package main

import "fmt"

func main(){
	var a, b int
	fmt.Scan(&a,&b)
	for a != b{
		if a > b{
			urutAngka(&a,&b)
		}
		fmt.Println(a,b)
		fmt.Scan(&a,&b)
	}
}
func urutAngka(a,b *int){
	var temp int
	temp = *a
	*a = *b
	*b = temp
}