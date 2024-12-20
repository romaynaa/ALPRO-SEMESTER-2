package main

import "fmt"

func lowToUpper(kar byte) byte{
	return kar - 32
}
func main(){
	
	inputChar := byte('a')
	outputChar := lowToUpper(inputChar)
	fmt.Printf("%c => %c\n", inputChar, outputChar)
}