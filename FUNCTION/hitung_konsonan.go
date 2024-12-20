package main

import "fmt"

func consonant(kar byte, konsonan bool)bool{
	konsonan = kar != 'a' && kar != 'i' && kar != 'u' && kar != 'e' && kar != 'o' && kar != 'A' && kar != 'I' && kar != 'U' && kar != 'E' && kar != 'O'
	return konsonan
}

func main(){
	var kata string
	var kons int
	fmt.Scan(&kata)

	for kata != '.'{
		kons++ 
		consonant(kons, kata)
		fmt.Println(kons)
	}

}