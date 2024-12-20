package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(hitungPangkat(n))

}
func hitungPangkat(n int)int{
	if n == 1{
		return 2
	}else{
		return 2 * hitungPangkat(n-1)
	}
}