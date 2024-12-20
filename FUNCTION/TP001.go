package main

import "fmt"

func lowToUpper(kar byte) byte{
	hasil := kar - 32
	return hasil
}
func main(){
	var a byte
	fmt.Scanf("%c", &a)
	fmt.Printf("%c", lowToUpper(a))
}