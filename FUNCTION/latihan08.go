package main

import "fmt"

func main(){
	var a,b int
	fmt.Scan(&a,&b)
	fmt.Print(hitungFaktorial(a)," ")
	fmt.Print(hitungFaktorial(b)," ")
	if a <= b{
		fmt.Print(hitungPermutasi(b,a))
	}else{
		fmt.Print(hitungPermutasi(a,b))
	}

}
func hitungFaktorial(n int)int{
	var fact int = 1
	for i := n; i >= 1; i--{
		fact = fact * i
	}
	return fact
}
func hitungPermutasi(a, b int)float64{
	var hasil float64
	hasil = float64(hitungFaktorial(a))/float64(hitungFaktorial(a-b))
	return hasil
}