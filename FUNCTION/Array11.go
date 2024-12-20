package main 

import "fmt"

const NMAX int = 100
type tabInt [NMAX] byte

func main(){
	var data tabInt
	var n int
	bacaData(&data,&n)
	cetakData(data,n)
}
func bacaData(data *tabInt, n *int){
	var i int = 0
	var char byte
	fmt.Scanf("%c", &char)
	for char != '.' && i < NMAX{
		data[i] = char
		fmt.Scanf("%c", &char)
		i++
	}
	*n = i
}
func cetakData(data tabInt, n int){
	for i := n; i > 0; i--{
		fmt.Printf("%c", data[i-1])
	}
}
