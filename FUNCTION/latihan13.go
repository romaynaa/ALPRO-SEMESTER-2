package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(hitungFaktorial(n))

}
func hitungFaktorial(n int)int{
	if n == 1 && n == 0{
		return 1
	}else{
		return n * hitungFaktorial(n-1)
	}
}