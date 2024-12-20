/*
program utama
kamus
	type pemain <
	nama, no_punggung string
	gol : integer
	jenis_kelamin : char
	>
	bil : integer
	pb : pemain
algoritma
	bil <- 1
	pb.nama <- "Shin Tae yong"
	pb.no_punggung <- "07"
	pb.gol <- 2
	pb.jenis_kelamin <- 'l'
output(pb.nama, pb.no_punggung)
output(pb.gol)
output(pb.jenis_kelamin)
*/
package main

import "fmt"

func main(){
	type pemain struct{
		nama, no_punggung string
		gol int
		jenis_kelamin byte
	}
	var bil int = 1
	var pb pemain

	pb.nama = "Shin Tae Yong"
	pb.no_punggung = "07"
	pb.gol = 2
	pb.jenis_kelamin = 'l'
	
	fmt.Print(bil, " ", pb.nama, " ", pb.no_punggung," ", pb.gol, " ")
	fmt.Printf(" %c\n", pb.jenis_kelamin)
}
