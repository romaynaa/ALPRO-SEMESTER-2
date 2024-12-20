package main

import "fmt"

func lowToUpper(kar byte)byte{
	return kar - 32
}
func main(){
	var huruf byte
	fmt.Scanf("%c",&huruf)
	fmt.Printf("%c",lowToUpper(huruf))
}