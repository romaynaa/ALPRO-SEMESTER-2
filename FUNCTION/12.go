package main

import "fmt"

func pecahDigit(ID int, no, lokasi *int, status *bool){
	*no = ID  / 1000
	*lokasi = (ID / 10) % 100
	*status = ID % 1000 == 1
}
func main(){
	var ID, no, lokasi int
	var status bool
	fmt.Scan(&ID)
	fmt.Println("No kelompok :",&no)
	fmt.Println("Lokasi ruangan :",&lokasi)
	fmt.Println("")
}