package main

import "fmt"

func upToLower(kar byte)byte{
	return kar + 32
}
func main(){
	var kar byte
	fmt.Scanf("%c", &kar)
	fmt.Printf("%c", upToLower(kar))
}