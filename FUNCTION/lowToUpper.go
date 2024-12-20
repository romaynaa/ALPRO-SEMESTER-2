package main

import "fmt"

func lowToUpper(kar byte) byte {
	hasil := kar -32
	return hasil
}
func main(){
	var huruf byte
	fmt.Scanf("%c", &huruf)
	fmt.Printf("%c", lowToUpper(huruf))
}