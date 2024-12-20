package main

import "fmt"

const NMAX int = 10
type tabInt[NMAX] byte

func main(){
	var info tabInt
	var ndata int
	bacaArray(&info, &ndata)
	cetakArray(info, ndata)
}
func bacaArray(info *tabInt, ndata *int){
	var kata byte
	var i  int = 0
	for kata != '.' && i < NMAX{
		info[i] = kata
		i += 1
		fmt.Scanf("%c", &kata)
	}
	*ndata = i
}
func cetakArray(info tabInt, ndata int){
	for i := ndata; i >= 0; i--{
		fmt.Printf("%c", info[i])
	}
	/*for ndata > 0{
		fmt.Printf("%c", info[ndata-1])
		ndata = ndata - 1
	}*/
}