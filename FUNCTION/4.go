package main

import "fmt"

func lowToUpper(kar byte)byte{
	return kar - 32
}
func main(){
	var kar byte
	fmt.Scanf("%c",&kar)
	fmt.Printf("%c", lowToUpper(kar))
}